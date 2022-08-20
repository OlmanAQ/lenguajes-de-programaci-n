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

func Stock(cal calzado, l *[]calzado) {
	for i := 0; i < 10; i++ {
		*l = append(*l, cal)
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
	var Adoc calzado
	var Guicci calzado

	lnike := []calzado{}
	ladidas := []calzado{}
	lUnderArmour := []calzado{}
	lCat := []calzado{}
	lAdoc := []calzado{}
	lGuicci := []calzado{}

	Stock(crearCalzado(34, 1250, "Nike", Nike), &lnike)
	Stock(crearCalzado(35, 3155, "Adidas", Adidas), &ladidas)
	Stock(crearCalzado(36, 2802, "UnderArmour", UnderArmour), &lUnderArmour)
	Stock(crearCalzado(37, 740, "Cat", Cat), &lCat)
	Stock(crearCalzado(38, 895, "Adoc", Adoc), &lAdoc)
	Stock(crearCalzado(43, 2121, "Guicci", Guicci), &lGuicci)

	fmt.Println("\nStock de Nike: ", lnike)
	fmt.Println("Tamaño de la lista: ", len(lnike))

	for i := 0; i < 10; i++ {
		vender(&lnike)
	}
	fmt.Println("\nDespues de 10 ventas: ", lnike)
	fmt.Println("Tamaño de la lista: ", len(lnike))

}
