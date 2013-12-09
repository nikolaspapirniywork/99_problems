module P2 where

{-
        (*) Find the last but one element of a list.

        (Note that the Lisp transcription of this problem is incorrect.)
        
        Example in Haskell:
        
        Prelude> myButLast [1,2,3,4]
        3
        Prelude> myButLast ['a'..'z']
        'y'
-}

myButLast :: [a] -> a
myButLast [] = error "No elements"
myButLast (_:[]) = error "Only one element"
myButLast (_:y:[]) = y
myButLast (_:xs) = myButLast xs

myButLast' :: [a] -> a
myButLast' list = case list of [] -> error "No elements"
                               (_:[]) -> error "Only one element"
                               (_:y:[]) -> y
                               (_:xs) -> myButLast' xs 
       
myButLast'' :: [a] -> a
myButLast'' [] = error "No elements"
myButLast'' (_:[]) = error "Only one element"
myButLast'' (x:xs) = foldr (\el _ -> el) x xs 


myButLast''' :: [a] -> a
myButLast''' [] = error "No elements"
myButLast''' (_:[]) = error "Only one element"
myButLast''' list = list !! (length list - 2)

myButLast'''' :: [a] -> a
myButLast'''' [] = error "No elements"
myButLast'''' (_:[]) = error "Only one element"
myButLast'''' list  = head $ tail $ reverse list


