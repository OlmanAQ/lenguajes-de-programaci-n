module Aplanacion where 

import Data.List

--  List flattening using map
--  map :: (a -> b) -> [a] -> [b]
--  concat :: [[a]] -> [a]
--  concatMap :: (a -> [b]) -> [a] -> [b]

aplanar :: [[a]] -> [a]
aplanar = concatMap id






main :: IO ()
main = do
    print $ aplanar [[1,2,3],[4,5,6],[7,8,9]]

   


