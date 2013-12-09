module P3 where

{-
        (*) Find the K'th element of a list. The first element in the list is number 1.
        
        Example:
        
        * (element-at '(a b c d e) 3)
        c
        Example in Haskell:
        
        Prelude> elementAt [1,2,3] 2
        2
        Prelude> elementAt "haskell" 5
        'e'
-}

elementAt :: [b] -> Int -> b
elementAt [] _ = error "No elements"
elementAt (x:xs) el
                | el > (length xs) + 1 || el <= 0 = error "Out of bounds"
                | el > 1 = elementAt xs $ el - 1
                | otherwise = x
                
elementAt' :: [b] -> Int -> b
elementAt' [] _ = error "No elements"
elementAt' list el
            | el > (length list) || el <= 0 = error "Out of bounds"
            | otherwise = list !! (el - 1)
            
           
elementAt'' :: [b] -> Int -> b
elementAt'' [] _ = error "No elements"
elementAt'' list el
            | el > (length list) || el <= 0 = error "Out of bounds"
            | otherwise = head $ drop (el - 1) list          
            
elementAt''' :: [b] -> Int -> b
elementAt''' [] _ = error "No elements"
elementAt''' list el
            | el > (length list) || el <= 0 = error "Out of bounds"
            | otherwise = head $ reverse $ take el list   
            
            
            