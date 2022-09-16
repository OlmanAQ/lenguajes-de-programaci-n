# Ejemplos2
calcularMontoVenta nombreProducto cantidadProductos listaProductos=
    concat (map montos listaProductos) 
    where
        montos prod = 
            if ((nombre prod) == nombreProducto) then
                [(cantidad prod) * (precio prod)]
            else
                []


