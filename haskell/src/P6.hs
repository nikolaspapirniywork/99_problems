module P6 where

{-
        6 Problem 6
        (*) Find out whether a list is a palindrome. A palindrome can be read forward or backward; e.g. (x a m a x).
        
        Example in Haskell:
        
        *Main> isPalindrome [1,2,3]
        False
        *Main> isPalindrome "madamimadam"
        True
        *Main> isPalindrome [1,2,4,8,16,8,4,2,1]
        True
-}

isPalindrome :: (Eq a) => [a] -> Bool
isPalindrome list = 
                if(isListEven) then 
                        (leftPart) == (rightPartEvenLength)
                else 
                        (leftPart) == (rightPartOddLength)  
                where 
                      isListEven = even $ length list
                      leftPart = fst splittedList 
                      splittedList = splitAt ((length list) `div` 2) list
                      rightPartEvenLength = reverse (snd splittedList)
                      rightPartOddLength = (reverse $ tail $ snd splittedList) -- since it breaks like [1,2] [3,2,1]
                      
isPalindrome' :: (Eq a) => [a] -> Bool
isPalindrome' xs = and $ zipWith (==) xs (reverse xs)                 

isPalindrome'' :: (Eq a) => [a] -> Bool
isPalindrome'' xs = foldr1 (&&) (zipWith (==) xs (reverse xs))

isPalindrome''' :: (Eq a) => [a] -> Bool
isPalindrome''' list = foldl (\acc (x,y) -> (x == y) && acc) True zippedList                    
                        where zippedList = list `zip` (reverse list)
                        
isPalindrome'''' :: (Eq a) => [a] -> Bool
isPalindrome'''' [] = True
isPalindrome'''' [_] = True
isPalindrome'''' x = (head x) == (last x) && (isPalindrome'''' $ init $ tail x) 






                                 