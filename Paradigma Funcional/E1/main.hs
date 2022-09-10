module Main where

import Data.List

--Haciendo uso de la función filter, implemente una función que a partir de una lista de cadenas de parámetro, filtre aquellas que contengan una subcadena que el usuario indique en otro argumento




sub_cadenas :: String -> [String] -> [String]
sub_cadenas str list = filter (isInfixOf str) list


main :: IO ()
main = do
    print $ sub_cadenas "la" ["la casa", "el perro", "pintando la cerca", "la casa de la esquina"]