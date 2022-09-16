module Facturas
    ( main
    ) where

data Producto = Producto { nombre :: String
                    , cantidad :: Int
                    , precio :: Float
                     } deriving Show


listarExistenciaMinima minimo listaProductos =
    filter minimos listaProductos 
    where
        minimos prod = ((cantidad prod) <= minimo) 




--calcularMontoVenta listaProductos = sum (map precio listaProductos)

calcularMontoVenta listaProductos = 
    foldl (\acum prod -> acum + (precio prod)) 0 listaProductos


--calcularMontoVenta con impuesto 13%
calcularMontoVentaImpuesto listaProductos = 
    foldl (\acum prod -> acum + (precio prod)) 0 listaProductos * 0.13


main :: IO ()
main = do
    let listaProductos=[
            (Producto "arroz" 15 2500),
            (Producto "frijoles" 4 2000),
            (Producto "leche" 8 1200),
            (Producto "cafe" 12 4500),
            (Producto "azucar" 20 1000),
            (Producto "sal" 30 500)]
    --print(listarExistenciaMinima 10 listaProductos)
    print(calcularMontoVenta listaProductos)

    
    montoVenta <- return (calcularMontoVenta listaProductos)
    print(montoVenta)
    impuesto <- return (calcularMontoVentaImpuesto listaProductos)
    print(impuesto)
    print(montoVenta + impuesto)


