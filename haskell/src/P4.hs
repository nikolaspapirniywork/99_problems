module P4 where

{-
        4 Problem 4
        (*) Find the number of elements of a list.
        
        Example in Haskell:
        
        Prelude> myLength [123, 456, 789]
        3
        Prelude> myLength "Hello, world!"
        13 
-}

myLength :: [a] -> Int
myLength [] = 0
myLength (_:xs) = 1 + myLength xs

myLength' :: (Num a) => [a] -> Int
myLength' [] = 0
myLength' list = foldr (\_ acc -> 1 + acc) 0 list

myLength'' :: (Num a) => [a] -> Int
myLength'' [] = 0
myLength'' list = foldl (\acc _ -> 1 + acc) 0 list