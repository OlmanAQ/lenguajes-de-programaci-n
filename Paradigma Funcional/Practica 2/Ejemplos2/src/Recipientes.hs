module Recipientes
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

vaciar_aa recipiente =
    0:(tail recipiente)
vaciar_bb recipiente =
    (head recipiente):[0]
llenar_aa recipiente =
    3:(tail recipiente)
llenar_bb recipiente =
    (head recipiente):[5]
pasar_ab recipiente =
    if (|| ((head recipiente) == 0)) ((head (tail recipiente)) == 5) then recipiente
    else pasar_ab (((head recipiente)-1):[((head (tail recipiente))+1)])
pasar_ba recipiente =
    if (|| ((head (tail recipiente)) == 0)) ((head recipiente) == 3) then recipiente
    else pasar_ba (((head recipiente)+1):[((head (tail recipiente))-1)])

vecinos recipiente =
    [(vaciar_aa recipiente),
    (vaciar_bb recipiente),
    (llenar_aa recipiente),
    (llenar_bb recipiente),
    (pasar_ab recipiente),
    (pasar_ba recipiente)]

extender ruta =
    remove_if (==[])
    (map (\x -> if (miembro x ruta) then []
                       else x:ruta )
        (vecinos (head ruta)))

prof recipiente objetivo =
    prof_aux [[recipiente]] objetivo

prof_aux ruta objetivo =
    if ruta==[] then []
    else if (solucion (head (head ruta)) objetivo) then
        reverse (head ruta)
        --reverse (head ruta):prof_aux ((extender (head ruta))++(tail ruta)) objetivo
    else
        prof_aux ((extender (head ruta))++(tail ruta)) objetivo

solucion recipiente objetivo =
    if (recipiente == objetivo) then
        True
    else
        False

--Implemente una función que encuentre el camino más corto en la búsqueda en profundidad basándose en la sumatoria de pesos
--de los nodos visitados. La función debe recibir como parámetros el estado inicial y el estado objetivo y debe devolver
--una lista de estados que representan el camino más corto desde el estado inicial hasta el estado objetivo.




main :: IO ()
main = do
    print(prof [2,2] [2,0])