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




--calcularMontoVenta listaProductos

calcularMontoVenta listaProductos = 
    sum (map precio listaProductos)




--calcularMontoVenta con impuesto 13%
calcularMontoVentaImpuesto listaProductos =
    sum (map precioImpuesto listaProductos)
    where
        precioImpuesto prod = (precio prod) * 0.13
    


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
    

    
    montoVenta <- return (calcularMontoVenta listaProductos)
    print("Monto de la venta: " ++ show(montoVenta))

    impuesto <- return (calcularMontoVentaImpuesto listaProductos)
    print("Impuesto: " ++ show(impuesto))

    montoTotal <- return (montoVenta + impuesto)
    print("Monto total: " ++ show(montoTotal))



