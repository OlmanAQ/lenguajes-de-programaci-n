package main

import (
	"fmt"
)

type calzado struct {
	marca  string
	precio int
	talla  int
	stock  int
}

type listaCalzado []calzado

var lCalzado listaCalzado

func (l *listaCalzado) venderCalzado(marca string, talla int, cant int) {
	var cal = l.buscarProducto(marca, talla)
	if cal != -1 && cant > 0 {
		if (*l)[cal].stock >= cant {
			(*l)[cal].stock = (*l)[cal].stock - cant
			if (*l)[cal].stock == 0 {
				lCalzado = append(lCalzado[:cal], lCalzado[cal+1:]...)
			}
		} else {
			fmt.Println("No se puede mayor cantidad de calzado de esta talla que los que hay en stock")
		}
	}
}

func stock() {
	lCalzado.agregarCalzado("Nike", 40, 15, 25000)
	lCalzado.agregarCalzado("Adidas", 37, 4, 20000)
	lCalzado.agregarCalzado("Under Armour", 39, 8, 12000)
	lCalzado.agregarCalzado("Cat", 41, 12, 45000)
	lCalzado.agregarCalzado("Guicci", 42, 6, 75000)
	lCalzado.agregarCalzado("Reebok", 40, 6, 25000)
	lCalzado.agregarCalzado("New Balance", 40, 6, 45000)
	lCalzado.agregarCalzado("Hi-Tec", 40, 7, 45000)
	lCalzado.agregarCalzado("Nike", 41, 15, 35000)
	lCalzado.agregarCalzado("Adidas", 40, 4, 20000)

}

func (l *listaCalzado) agregarCalzado(nombre string, talla int, cantidad int, precio int) {
	var cal = l.buscarProducto(nombre, talla)
	if (cal) != -1 {
		(*l)[cal].stock = (*l)[cal].stock + cantidad
		if ((*l)[cal].precio) != precio {
			(*l)[cal].precio = precio
		}
	} else {
		lCalzado = append(lCalzado, calzado{marca: nombre, talla: talla, stock: cantidad, precio: precio})
	}

	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente

}

func (l *listaCalzado) buscarProducto(nombre string, talla int) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].marca == nombre && (*l)[i].talla == talla {
			result = i
		}
	}
	return result
}

func main() {

	stock()
	fmt.Println("\nStock de calzado: ", lCalzado)
	fmt.Println("Tamaño de la lista: ", len(lCalzado))

	lCalzado.venderCalzado("Nike", 40, 15)
	fmt.Println("\nDespues de 15 ventas de nike: ", lCalzado)
	fmt.Println("Tamaño de la lista: ", len(lCalzado))

}
