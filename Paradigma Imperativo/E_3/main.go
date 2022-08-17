package main

import (
	"fmt"
)

func rotarArreglo(arr *[8]string, dirreccion string, rotaciones int) {

	for i := range arr {

	}

}

func main() {

	arr := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	dirreccion := "Derecha"
	var rotar = 5

	fmt.Println("Arreglo: ", arr)
	fmt.Println("Direccion: ", dirreccion)
	fmt.Println("Rotacciones: ", rotar)

	rotarArreglo(&arr, dirreccion, rotar)

}
