module Aplanacion where 

--  Aplanacion de una lista anidas de listas 
-- Ejemplo: 
-- aplanar [[1,[1,2],[3,4]], [[5,6],[7,8]]]
-- [1,1,2,3,4,5,6,7,8]


aplanar :: [[a]] -> [a]
aplanar [] = []
aplanar (x:xs) = x ++ aplanar xs


main :: IO ()
main = do
    print (aplanar [[[1,2],[3,4]], [[5,6],[7,8]]])  
