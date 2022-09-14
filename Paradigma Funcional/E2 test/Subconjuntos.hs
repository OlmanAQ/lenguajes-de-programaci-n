
module Subconjuntos
    ( main
    ) where


--Esta función recibe dos listas y debe retornar True cuando el primer argumento es subconjunto completo del segundo y False en caso contrario.
--Ejemplos:
--sub_conjunto [1,2,3] [1,2,3,4,5] == True
--sub_conjunto [1,2,3] [1,2,4,5] == False

sub_conjunto :: [Int] -> [Int] -> Bool
sub_conjunto [] _ = True
sub_conjunto (i:aux) y = elem i y && sub_conjunto aux y




main :: IO ()
main = do
    print $ sub_conjunto [1,2,3] [1,2,3,4,5]
    print $ sub_conjunto [1,2,3] [1,2,4,5]
    
