module Grafos
    ( main
    ) where

import Data.Ord (comparing)
import Data.List (sortBy)

combine :: [a] -> [a] -> [[a]]
combine p1 p2 = [[a,b] | (a, b) <- zip p1 p2]

transformar :: [String] -> [(String, Integer)]
transformar [] = []
transformar (k:v:t) = (k, read v) : transformar t

memb element lista =
    or (map comp lista)
    
    where
        comp x = (x == element)

remove_if fun lista =
    concat (map (\x -> if (fun x) then []
                       else [x] )
                lista)

usar_vertices nodo grafo = do
    let res =  concat (map (\x -> if (head (head x)) == nodo then [x]
                            else [] ) 
                            grafo)
    if (res == []) then []
    else head (tail (head res))

extender ruta grafo =
    remove_if (==[])
    (map (\x -> if (memb x ruta) then []
                       else x:ruta )
        (orden_pesos (head ruta) grafo))

prof ini fin grafo =
    prof_aux fin [[ini]] grafo

ord :: [[a]] -> [a]
ord []=[]
ord [y] = y
ord (x:y:list)
 |length x > length y = ord (y:list)            
 |otherwise = ord (x:list)
 
usar_pesos nodo grafo = do

    let res =  concat (map (\x -> if (head(head x)) == nodo then [tail x]
                            else [] ) 
                            grafo)
    if (res == []) then []
    else head (tail (head res))

prof_aux fin routes grafo =

    if routes==[] then []
    else if (fin == (head (head routes))) then
        reverse (head routes):prof_aux fin ((extender (head routes) grafo)++(tail routes)) grafo
    else
        prof_aux fin ((extender (head routes) grafo)++(tail routes)) grafo
    

orden_pesos nodo grafo = do

    let {combinar = combine (usar_vertices nodo grafo) (usar_pesos nodo grafo); concatenar = concat combinar; 

    to_StringInt = transformar concatenar; ordenarConcat = sortBy (comparing (\(x,y) -> abs y)) to_StringInt; 

    del_pesos = to_StringInt >>= \(x,y) -> [x]} in del_pesos

-- Funcion Main

main :: IO ()
main = do

    let grafo=[ [["i"],["a","b"], ["3","6"]],     
                [["a"],["i","c","d"], ["3","6","5"]], 
                [["b"],["i","c","d"], ["3","6","5"]],  
                [["c"],["a","b","x"], ["3","6","5"]],  
                [["d"],["a","b","f"], ["3","6","5"]], 
                [["f"],["d"], ["3"]],          
                [["x"],["c"], ["3"]]           
                 ]   
    
    let lista=prof "i" "f" grafo

    let new = ord lista

    print("El camino mas corto encontrado es: ")
    print(new)