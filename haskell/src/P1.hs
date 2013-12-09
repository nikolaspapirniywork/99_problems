module P1 where
{-
        (*) Find the last element of a list.
        
        (Note that the Lisp transcription of this problem is incorrect.)
        
        Example in Haskell:
        
        Prelude> myLast [1,2,3,4]
        4
        Prelude> myLast ['x','y','z']
        'z'
-}

lastElement :: [a] -> a
lastElement [] = error "Can't take last element from empty list"
lastElement (x:[]) = x
lastElement (_:xs) = lastElement xs
        
lastElement' :: [a] -> a
lastElement' [] = error "Can't take last element from empty list" 
lastElement' (x:xs) = foldl (\_ el -> el) x xs

lastElement'' :: [a] -> a
lastElement'' list = head $ reverse list

