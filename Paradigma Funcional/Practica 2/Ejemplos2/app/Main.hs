module Main (main) where

--import Lib

--miembro1 :: Int -> [Int] -> Bool
miembro1 e l =
    if l == [] then
        False
    else if e == (head l) then
        True
    else 
        miembro1 e (tail l)

miembro2 :: (Eq a) => a -> [a] -> Bool
miembro2 _ [] = False
miembro2 e (x:xs) 
    | e == x = True
    | otherwise = miembro2 e xs


--menor :: [a] -> Either a [a]
menor (l)
    | l == [] =  -1 --averiguar como devolver dos tipos diferentes y eventualmente como imprimir resultados de diferentes tipos
    | otherwise = menor_aux (head l) (tail l)

--menor_aux :: (Ord a, Num a) => a -> [a] -> a
menor_aux min [] = min
menor_aux min (h:t)
    | min < h = menor_aux min t
    | min >= h = menor_aux h t

--eliminar :: (Eq a) => a -> [a] -> [a]
eliminar _ [] = [] 
eliminar elem (h:t)  
    | elem == h = eliminar elem t
    | otherwise = h:(eliminar elem t)

eliminar_map elem lista =
    concat (map compare' lista)
    where
        compare' e = 
            if (e == elem) then 
                []
            else 
                [e] 

main :: IO ()
main = do 
    let lista = [3,4,1,2,3,5,3]
    let lista2 = [[],[1],[],[2]]
    print(miembro1 4 lista)
    print(miembro1 6 lista)

    print(miembro2 4 lista)
    print(miembro2 6 lista)
    print(miembro2 "mundo" ["hola","mundo","hola"])

    print(menor lista)
    
    print(eliminar 3 lista)
    print(eliminar "mundo" ["hola","mundo","hola"])
    print lista

    print(eliminar_map 3 lista)
    print lista

