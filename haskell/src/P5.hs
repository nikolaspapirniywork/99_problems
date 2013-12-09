module P5 where

{-
        (*) Reverse a list.
        
        Example in Haskell:
        
        Prelude> myReverse "A man, a plan, a canal, panama!"
        "!amanap ,lanac a ,nalp a ,nam A"
        Prelude> myReverse [1,2,3,4]
        [4,3,2,1]
-}

myReverse :: [a] -> [a]
myReverse [] = []
myReverse [x] = [x]
myReverse (x:xs) = myReverse xs ++ [x]
 
myReverse' :: [a] -> [a]
myReverse' [] = []
myReverse' [x] = [x]
myReverse' list = foldl (\ acc x -> x : acc) [] list 

myReverse'' :: [a] -> [a]
myReverse'' [] = []
myReverse'' [x] = [x]
myReverse'' list = foldr (\ x acc -> acc ++ [x]) [] list

myReverse''' :: [a] -> [a]
myReverse''' [] = []
myReverse''' [x] = [x]
myReverse''' list = foldl (flip (:)) [] list 

myReverse'''' :: [a] -> [a]
myReverse'''' list = myReverse_acc list [] 
        where
                myReverse_acc [] reversed = reversed
                myReverse_acc (x:xs) reversed = myReverse_acc xs (x:reversed)
                 
             


