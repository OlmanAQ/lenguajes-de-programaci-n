module Grafos
    ( main
    ) where

--miembro 
miembro elem lista =
    or (map comparacion lista)
    --any (==True) (map comparacion lista)
    where
        comparacion x = (x == elem)

--remove_if
remove_if fun lista =
    concat (map (\x -> if (fun x) then []
                       else [x] )
                lista)

vecinos nodo grafo = do
    let resultado =  concat (map (\x -> if (head (head x)) == nodo then [x]
                            else [] ) 
                            grafo)
    if (resultado == []) then []
    else head (tail (head resultado))

extender ruta grafo =
    remove_if (==[])
    (map (\x -> if (miembro x ruta) then []
                       else x:ruta )
        (vecinos (head ruta) grafo))

prof ini fin grafo =
    prof_aux fin [[ini]] grafo

prof_aux fin rutas grafo =
    if rutas==[] then []
    else if (fin == (head (head rutas))) then
        reverse (head rutas)
        --reverse (head rutas):prof_aux fin ((extender (head rutas) grafo)++(tail rutas)) grafo
    else
        prof_aux fin ((extender (head rutas) grafo)++(tail rutas)) grafo

--Programa que resuelva el problema de encontrar la o las rutas de un laberinto con datos quemados en el programa.





main :: IO ()
main = do
    let grafo=[ [["i"],["a","b"]],      --        a ---- c ---- x
                [["a"],["i","c","d"]],  --      /   \  /
                [["b"],["i","c","d"]],  --     /     \/
                [["c"],["a","b","x"]],  --   i       X
                [["d"],["a","b","f"]],  --     \    / \
                [["f"],["d"]],          --      \  /   \
                [["x"],["c"]]           --        b ---- d ---- f
                 ]           

    print(prof "i" "f" grafo)


    