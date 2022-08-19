package main

import "fmt"

type calzado struct {
	marca  string
	precio int
	talla  int
}

func vender(l *[]calzado) {
	aux := *l
	if len(aux) == 0 {
		println("Calzado Agotado")
		aux = append(aux[:0], aux[1:]...)
	} else {
		aux = append(aux[:0], aux[1:]...)
		*l = aux
	}
}

func Stock(cal calzado, lista *[]calzado) {
	for i := 0; i < 10; i++ {
		*lista = append(*lista, cal)
	}
}

func crearCalzado(t int, p int, m string, zapato calzado) calzado {

	zapato.talla = t
	zapato.precio = p
	zapato.marca = m

	return zapato
}

func main() {

	var Nike calzado
	var Adidas calzado
	var UnderArmour calzado
	var Cat calzado
	var Magnanni calzado
	var Bally calzado
	var EdwardGreen calzado
	var JohnLobb calzado
	var Buscemi calzado
	var Guidi calzado

	// lC = lista Calzado[NombreCalzado]

	lCGeorgeCleverley := []calzado{}
	lCoADiciannoveventitre := []calzado{}
	lCSebastianTarek := []calzado{}
	lCSantoni := []calzado{}
	lCMagnanni := []calzado{}
	lCBally := []calzado{}
	lCEdwardGreen := []calzado{}
	lCJohnLobb := []calzado{}
	lCBuscemi := []calzado{}
	lCGuidi := []calzado{}

	Stock(crearCalzado(34, 1250, "George Cleverley", Nike), &lCGeorgeCleverley)
	Stock(crearCalzado(35, 3155, "A Diciannoveventitre", Adidas), &lCoADiciannoveventitre)
	Stock(crearCalzado(36, 2802, "Sebastian Tarek", UnderArmour), &lCSebastianTarek)
	Stock(crearCalzado(37, 740, "Santoni", Cat), &lCSantoni)
	Stock(crearCalzado(38, 895, "Magnanni", Magnanni), &lCMagnanni)
	Stock(crearCalzado(39, 650, "Bally", Bally), &lCBally)
	Stock(crearCalzado(40, 2000, "Edward Green", EdwardGreen), &lCEdwardGreen)
	Stock(crearCalzado(41, 1625, "John Lobb", JohnLobb), &lCJohnLobb)
	Stock(crearCalzado(42, 1887, "Buscemi", Buscemi), &lCBuscemi)
	Stock(crearCalzado(43, 2121, "Guidi", Guidi), &lCGuidi)

	fmt.Println("\nLista de zapatos inicial: ", lCGeorgeCleverley)
	fmt.Println("Tamaño de la lista: ", len(lCGeorgeCleverley))

	for i := 0; i < 10; i++ {
		vender(&lCGeorgeCleverley)
	}
	fmt.Println("\nLista de zapatos despues de 10 ventas: ", lCGeorgeCleverley)
	fmt.Println("Tamaño de la lista: ", len(lCGeorgeCleverley))

}
