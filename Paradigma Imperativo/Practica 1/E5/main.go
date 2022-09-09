package main

import (
	"fmt"
)

func main() {

	//var exp1 string = "456*+"
	//var exp2 string = "78+32+/"
	var exp3 string = "28+5-"

	pilaSliceOperadores := []int{}

	reslFinal := decodificarExpresionInfija(exp3, pilaSliceOperadores)
	//fmt.Println("\nExpresion numerica, exp1: ", exp1, " resultado: (5*6)+4 = 34")
	//fmt.Println("\nExpresion numerica, exp2: ", exp2, " resultado: (7+8)/(3+2) = 3")
	fmt.Println("\nExpresion numerica, exp3: ", exp3, " resultado: (2+8)-5 = 5")
	fmt.Println("\nResultado final: ", reslFinal)
}

func decodificarExpresionInfija(exp string, lista []int) int {

	for i := 0; i < len(exp); i++ {
		switch string(exp[i]) {
		case "+":
			result := operar(lista[(len(lista)-2)], lista[(len(lista)-1)], "+")
			listat := push(lista)
			lista = listat
			listat = push(lista)
			lista = listat
			lista = pop(lista, result)

		case "-":
			result := operar(lista[(len(lista)-2)], lista[(len(lista)-1)], "-")
			listat := push(lista)
			lista = listat
			listat = push(lista)
			lista = listat
			lista = pop(lista, result)

		case "/":
			result := operar(lista[(len(lista)-2)], lista[(len(lista)-1)], "/")
			listat := push(lista)
			lista = listat
			listat = push(lista)
			lista = listat
			lista = pop(lista, result)

		case "*":
			result := operar(lista[(len(lista)-2)], lista[(len(lista)-1)], "*")
			listat := push(lista)
			lista = listat
			listat = push(lista)
			lista = listat
			lista = pop(lista, result)

		default:
			var data int = int(exp[i] - 48)
			lista = pop(lista, data)
		}
	}

	return lista[0]
}

func push(lista []int) []int {
	listaNueva := lista[0:(len(lista) - 1)]
	return listaNueva
}
func pop(lista []int, num int) []int {
	lista = append(lista, num)
	return lista
}
func operar(num1 int, num2 int, operador string) int {

	switch operador {
	case "+":
		resultado := num1 + num2
		return resultado
	case "-":
		resultado := num1 - num2
		return resultado
	case "/":
		resultado := num1 / num2
		return resultado
	case "*":
		resultado := num1 * num2
		return resultado
	}

	return 0
}
