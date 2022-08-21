package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}

type listaProductos []producto

var lProductos listaProductos

var lProductosMinimos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	var prod = l.buscarProducto(nombre)
	if (prod) != -1 {
		(*l)[prod].cantidad = (*l)[prod].cantidad + cantidad
		if ((*l)[prod].precio) != precio {
			(*l)[prod].precio = precio
		}
	} else {
		lProductos = append(lProductos, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}

	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente

}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 && cant > 0 {
		if (*l)[prod].cantidad >= cant {
			(*l)[prod].cantidad = (*l)[prod].cantidad - cant
			if (*l)[prod].cantidad == 0 {
				lProductos = append(lProductos[:prod], lProductos[prod+1:]...)
			}
		} else {
			fmt.Println("No se puede vayor cantidad de productos que los que hay en existencia")
		}

		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista"
	}
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
}

func listarProductosMínimos(l *listaProductos) {
	for i := 0; i < len((*l)); i++ {
		if (*l)[i].cantidad <= existenciaMinima {
			lProductosMinimos = append(lProductosMinimos, (*l)[i])
		}
	}
}
func aumentarInventarioDeMinimos(l *listaProductos) {
	for producto := 0; producto < len((*l)); producto++ {
		if (*l)[producto].cantidad <= existenciaMinima {
			(*l)[producto].cantidad = existenciaMinima
		}
	}
}

func ordenarInventario(l *listaProductos, llave int) {
	switch llave {
	case 1:
		sort.Slice(*l, func(i, j int) bool { return (*l)[i].nombre < (*l)[j].nombre })
		fmt.Println("Orden por llave nombre")
	case 2:
		sort.Slice(*l, func(i, j int) bool { return (*l)[i].cantidad < (*l)[j].cantidad })
		fmt.Println("Orden por llave cantidad")
	case 3:
		sort.Slice(*l, func(i, j int) bool { return (*l)[i].precio < (*l)[j].precio })
		fmt.Println("\nOrden por llave precio")
	}
}

func main() {
	llenarDatos()

	fmt.Println("\nLista de Productos inicial: ", lProductos)

	listarProductosMínimos(&lProductos)
	fmt.Println("\nLista de Productos minimos: ", lProductosMinimos)

	aumentarInventarioDeMinimos(&lProductosMinimos)
	fmt.Println("\nLista de Productos con minimos constantes: ", lProductosMinimos)

	/*
		fmt.Println("Lista de productos inicial: ", lProductos)
		ordenarInventario(&lProductos, 1)

		fmt.Println(lProductos)
	*/

}
