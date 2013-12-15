module P8 where

import Data.List

{-

(**) Eliminate consecutive duplicates of list elements.

If a list contains repeated elements they should be replaced with a single copy of the element. 
The order of the elements should not be changed.

Example:

* (compress '(a a a a b c c a a d e e e e))
(A B C A D E)

-}

compress :: (Eq a) => [a] -> [a]
compress [] = []
compress list = foldr (\y acc -> if elem y acc then acc else y : acc) [] list  

compress' :: (Eq a) => [a] -> [a]
compress' (x:ys@(y:_))
    | x == y    = compress ys
    | otherwise = x : compress ys
compress' ys = ys

                                
compress'' :: (Eq a) => [a] -> [a]
compress'' = map head . group 