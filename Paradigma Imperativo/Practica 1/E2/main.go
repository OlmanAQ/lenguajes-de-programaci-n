package main

import "fmt"

func main() {

	fmt.Println()
	rombo(11)
}

func rombo(num int) {

	//Triangulo superior
	for row := 0; row < num; row += 2 {

		var noOfSpace int = num - row
		for space := 0; space <= noOfSpace; space++ {
			fmt.Print(" ")
		}
		noOfChar := row
		for column := 0; column <= noOfChar; column++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	//Parte inferior al triangulo
	number := (num - 2)
	for row := 0; row < number; row += 2 { // num= 5
		var noOfSpace int = 4 + row
		for space := 0; space < noOfSpace; space++ {
			fmt.Print(" ")
		}
		noOfChar := (num + 2) - noOfSpace
		for column := 0; column < noOfChar; column++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
