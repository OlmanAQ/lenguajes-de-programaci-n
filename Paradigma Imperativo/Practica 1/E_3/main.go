package main

import (
	"fmt"
)

func rotarArreglo(arr *[8]string, direccion string, rotar int) [8]string {

	for rot := 0; rot < rotar; rot++ {
		str1 := "Derecha"
		result := str1 == direccion
		if result == true {
			var aux string = arr[7]
			for i := len(arr) - 1; i > -1; i-- {
				if i != 0 {
					arr[i] = arr[i-1]
				} else {
					arr[0] = aux
				}

			}

		} else {
			var aux string = arr[0]
			for i := 0; i < len(arr); i++ {
				if i != 7 {
					arr[i] = arr[i+1]
				} else {
					arr[7] = aux
				}
			}
		}
	}
	return *arr
}

func main() {

	arr := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	dirreccion := "Derecha"
	var rotar = 3

	fmt.Println("Arreglo: ", arr)
	fmt.Println("Direccion: ", dirreccion)
	fmt.Println("Rotacciones: ", rotar)

	var newarr = rotarArreglo(&arr, dirreccion, rotar)
	fmt.Println("Arreglo rotado: ", newarr)

}
