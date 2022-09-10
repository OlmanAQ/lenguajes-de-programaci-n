module Main where

import Data.List

--Using the filter function, implement a function that, from a list of parameter strings, filters those that contain a substring that the user indicates in another argument.

--For example, if the user enters the string "a" and the list ["abc", "def", "ghi"], the result should be ["abc", "def"].


sub_cadenas :: String -> [String] -> [String]
sub_cadenas str list = filter (isInfixOf str) list


main :: IO ()
main = do
    print $ sub_cadenas "la" ["la casa", "el perro", "pintando la cerca", "la casa de la esquina"]