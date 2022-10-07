package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type asiento struct {
	codeAsiento string
	estado      string
	numAsiento  int
	categoria   string
	precios     int
	fila        int
}

type asientos []asiento

var listaAsientos asientos

func (a *asientos) agregarAsiento(code string, estado string, num int, categoria string, precios int, fila int) {
	idx := slices.IndexFunc(*a, func(p asiento) bool { return p.codeAsiento == code })
	if idx == -1 {
		*a = append(*a, asiento{code, estado, num, categoria, precios, fila})
	} else {
		fmt.Println("cant add existing product to the list")
	}
}

func main() {
	listaAsientos.agregarAsiento("A1", "disponible", 1, "primera", 100, 1)
	listaAsientos.agregarAsiento("A2", "disponible", 2, "primera", 100, 1)
	listaAsientos.agregarAsiento("A3", "disponible", 3, "primera", 100, 1)
	listaAsientos.agregarAsiento("A4", "disponible", 4, "primera", 100, 1)
	listaAsientos.agregarAsiento("A5", "disponible", 5, "primera", 100, 1)
	listaAsientos.agregarAsiento("A6", "disponible", 6, "primera", 100, 2)
	listaAsientos.agregarAsiento("A7", "disponible", 7, "primera", 100, 2)
	listaAsientos.agregarAsiento("A8", "disponible", 8, "primera", 100, 2)
	listaAsientos.agregarAsiento("A9", "disponible", 9, "primera", 100, 2)
	listaAsientos.agregarAsiento("A10", "disponible", 10, "primera", 100, 2)
	listaAsientos.agregarAsiento("A11", "disponible", 11, "primera", 100, 3)
	listaAsientos.agregarAsiento("A12", "disponible", 12, "primera", 100, 3)
	listaAsientos.agregarAsiento("A13", "disponible", 13, "primera", 100, 3)
	listaAsientos.agregarAsiento("A14", "disponible", 14, "primera", 100, 3)
	listaAsientos.agregarAsiento("A15", "disponible", 15, "primera", 100, 3)

	listaAsientos.agregarAsiento("B1", "disponible", 1, "primera", 100, 1)
	listaAsientos.agregarAsiento("B2", "disponible", 2, "primera", 100, 1)
	listaAsientos.agregarAsiento("B3", "disponible", 3, "primera", 100, 1)
	listaAsientos.agregarAsiento("B4", "disponible", 4, "primera", 100, 1)
	listaAsientos.agregarAsiento("B5", "disponible", 5, "primera", 100, 1)
	listaAsientos.agregarAsiento("B6", "disponible", 6, "primera", 100, 2)
	listaAsientos.agregarAsiento("B7", "disponible", 7, "primera", 100, 2)
	listaAsientos.agregarAsiento("B8", "disponible", 8, "primera", 100, 2)
	listaAsientos.agregarAsiento("B9", "disponible", 9, "primera", 100, 2)
	listaAsientos.agregarAsiento("B10", "disponible", 10, "primera", 100, 2)
	listaAsientos.agregarAsiento("B11", "disponible", 11, "primera", 100, 3)
	listaAsientos.agregarAsiento("B12", "disponible", 12, "primera", 100, 3)
	listaAsientos.agregarAsiento("B13", "disponible", 13, "primera", 100, 3)
	listaAsientos.agregarAsiento("B14", "disponible", 14, "primera", 100, 3)
	listaAsientos.agregarAsiento("B15", "disponible", 15, "primera", 100, 3)

	listaAsientos.agregarAsiento("C1", "disponible", 1, "primera", 100, 1)
	listaAsientos.agregarAsiento("C2", "disponible", 2, "primera", 100, 1)
	listaAsientos.agregarAsiento("C3", "disponible", 3, "primera", 100, 1)
	listaAsientos.agregarAsiento("C4", "disponible", 4, "primera", 100, 1)
	listaAsientos.agregarAsiento("C5", "disponible", 5, "primera", 100, 1)
	listaAsientos.agregarAsiento("C6", "disponible", 6, "primera", 100, 2)
	listaAsientos.agregarAsiento("C7", "disponible", 7, "primera", 100, 2)
	listaAsientos.agregarAsiento("C8", "disponible", 8, "primera", 100, 2)
	listaAsientos.agregarAsiento("C9", "disponible", 9, "primera", 100, 2)
	listaAsientos.agregarAsiento("C10", "disponible", 10, "primera", 100, 2)
	listaAsientos.agregarAsiento("C11", "disponible", 11, "primera", 100, 3)
	listaAsientos.agregarAsiento("C12", "disponible", 12, "primera", 100, 3)
	listaAsientos.agregarAsiento("C13", "disponible", 13, "primera", 100, 3)
	listaAsientos.agregarAsiento("C14", "disponible", 14, "primera", 100, 3)
	listaAsientos.agregarAsiento("C15", "disponible", 15, "primera", 100, 3)

	listaAsientos.agregarAsiento("A1", "disponible", 1, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("A2", "disponible", 2, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("A3", "disponible", 3, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("A4", "disponible", 4, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("A5", "disponible", 5, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("A6", "disponible", 6, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("A7", "disponible", 7, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("A8", "disponible", 8, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("A9", "disponible", 9, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("A10", "disponible", 10, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("A11", "disponible", 11, "Segunda", 100, 3)
	listaAsientos.agregarAsiento("A12", "disponible", 12, "Segunda", 100, 3)
	listaAsientos.agregarAsiento("A13", "disponible", 13, "Segunda", 100, 3)
	listaAsientos.agregarAsiento("A14", "disponible", 14, "Segunda", 100, 3)
	listaAsientos.agregarAsiento("A15", "disponible", 15, "Segunda", 100, 3)

	listaAsientos.agregarAsiento("B1", "disponible", 1, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("B2", "disponible", 2, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("B3", "disponible", 3, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("B4", "disponible", 4, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("B5", "disponible", 5, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("B6", "disponible", 6, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("B7", "disponible", 7, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("B8", "disponible", 8, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("B9", "disponible", 9, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("B10", "disponible", 10, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("B11", "disponible", 11, "Segunda", 100, 3)
	listaAsientos.agregarAsiento("B12", "disponible", 12, "Segunda", 100, 3)
	listaAsientos.agregarAsiento("B13", "disponible", 13, "Segunda", 100, 3)
	listaAsientos.agregarAsiento("B14", "disponible", 14, "Segunda", 100, 3)
	listaAsientos.agregarAsiento("B15", "disponible", 15, "Segunda", 100, 3)

	listaAsientos.agregarAsiento("C1", "disponible", 1, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("C2", "disponible", 2, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("C3", "disponible", 3, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("C4", "disponible", 4, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("C5", "disponible", 5, "Segunda", 100, 1)
	listaAsientos.agregarAsiento("C6", "disponible", 6, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("C7", "disponible", 7, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("C8", "disponible", 8, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("C9", "disponible", 9, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("C10", "disponible", 10, "Segunda", 100, 2)
	listaAsientos.agregarAsiento("C11", "disponible", 11, "Segunda", 100, 3)
	listaAsientos.agregarAsiento("C12", "disponible", 12, "Segunda", 100, 3)
	listaAsientos.agregarAsiento("C13", "disponible", 13, "Segunda", 100, 3)
	listaAsientos.agregarAsiento("C14", "disponible", 14, "Segunda", 100, 3)
	listaAsientos.agregarAsiento("C15", "disponible", 15, "Segunda", 100, 3)

	fmt.Println(listaAsientos)
}
