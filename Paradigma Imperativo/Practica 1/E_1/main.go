package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Hola mundo\nen lenguaje GO"

	fmt.Println("\nEl string: ", text)

	caracteres := len(text)

	saltos := strings.Count(text, "\n")

	var palabras int
	campos := strings.Fields(text)
	palabras = len(campos)

	fmt.Println("\nSaltos de linea: ", saltos)
	fmt.Println("Caracteres: ", caracteres)
	fmt.Println("Palabras: ", palabras)

}
