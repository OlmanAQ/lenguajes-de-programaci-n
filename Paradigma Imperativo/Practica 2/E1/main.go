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

type lProductosMinimos []producto

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	producto := producto{nombre: nombre, cantidad: cantidad, precio: precio}
	if len(*l) == 0 {
		*l = append(*l, producto)
	} else {
		for i := 0; i < len(*l); i++ {
			if (*l)[i].nombre == nombre {
				(*l)[i].cantidad += 1
				if (*l)[i].precio != precio {
					(*l)[i].precio = precio
				}
			} else {
				*l = append(*l, producto)
				break
			}
		}
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
		if (*l)[prod].cantidad != 0 {
			if (*l)[prod].cantidad > cant {
				(*l)[prod].cantidad = (*l)[prod].cantidad - cant
			} else if (*l)[prod].cantidad == cant {
				(*l) = append((*l)[:prod], (*l)[prod+1:]...)
				fmt.Println("Producto vendido se agotaron todas las existencias ")
			} else {
				fmt.Println("No se puede vender mayor cantidad de productos que los que hay en existencia")
				fmt.Println("Cantidad de productos en stock:", (*l)[prod].cantidad)
			}
		} else {
			fmt.Println("El producto se encuentra agotado.")
		}
	} else {
		fmt.Println("El producto no se encuentra en stock")
	}
}
func llenarDatos() {
	lProductos.agregarProducto("Arroz", 15, 2500)
	lProductos.agregarProducto("Frijoles", 4, 2000)
	lProductos.agregarProducto("Leche", 8, 1200)
	lProductos.agregarProducto("Café", 12, 4500)
	lProductos.agregarProducto("Azucar", 10, 1500)
	lProductos.agregarProducto("Pescado", 5, 3500)
	lProductos.agregarProducto("Queso", 20, 5500)
	lProductos.agregarProducto("Mantequilla", 7, 800)

}

func listarProdMínimos(l *listaProductos) []producto {
	var listaMinimos []producto
	for i := 0; i < len((*l)); i++ {
		if (*l)[i].cantidad <= existenciaMinima {
			listaMinimos = append(listaMinimos, (*l)[i])
		}
	}
	return listaMinimos
}
func aumentarInventarioDeMinimos(l *listaProductos) {
	for producto := 0; producto < len((*l)); producto++ {
		if (*l)[producto].cantidad <= 10 {
			(*l)[producto].cantidad = 10
		}
	}
}

// SORT SLICE OF STRUCTURES
func sortSliceOfStructures(l *listaProductos, atributo string) {
	switch atributo {
	case "nombre":
		sort.Slice(*l, func(i, j int) bool { return (*l)[i].nombre < (*l)[j].nombre })
		fmt.Println("\nOrdenado por el nombre:", *l)
	case "cantidad":
		sort.Slice(*l, func(i, j int) bool { return (*l)[i].cantidad < (*l)[j].cantidad })
		fmt.Println("\nOrdenado por la cantidad:", *l)
	case "precio":
		sort.Slice(*l, func(i, j int) bool { return (*l)[i].precio < (*l)[j].precio })
		fmt.Println("\nOrdenado por el precio:", *l)
	}
}

func main() {
	llenarDatos()
	/*
		fmt.Println("\nLista de productos inicial: ", lProductos, "\n")
		sortSliceOfStructures(&lProductos, "nombre")
	*/
	fmt.Println("\n------------------ lista inicial ------------------")
	fmt.Println("\nLista de Productos inicial: ", lProductos)

	fmt.Println("\n------------------ Lista de minimos ------------------")
	lProductosMinimos := listarProdMínimos(&lProductos)
	fmt.Println("\nLista de Productos minimos: ", lProductosMinimos)

	fmt.Println("\n---------- Aumento de cantidad a productos minimos en lista productos ----------")
	aumentarInventarioDeMinimos(&lProductos)
	fmt.Println("\nLista de Productos con minimos aumentados: ", lProductos)

}
