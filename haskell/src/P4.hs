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

myLength''' :: [a] -> Int
myLength''' list = mylist_acc 0 list
                where 
                        mylist_acc :: Int -> [a] -> Int
                        mylist_acc acc [] = acc
                        mylist_acc acc (_:xs) = mylist_acc (acc + 1) xs
                        
myLength'''' :: [a] -> Int
myLength'''' list =
                let
                        mylist_acc :: [a] -> Int -> Int
                        mylist_acc [] acc  = acc
                        mylist_acc (_:xs) acc = mylist_acc xs (acc + 1)
                in        
                        mylist_acc list 0             
                        
myLength''''' :: [a] -> Int 
myLength''''' list = sum $ map (\_ -> 1) list                      
                        
