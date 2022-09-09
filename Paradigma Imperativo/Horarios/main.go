package main

import "fmt"

type curso struct {
	nombre       string
	dia          int
	hora         bool
	aula         string
	lEstudiantes []estudiante
}

type estudiante struct {
	nombre string
	nota   int
}

type listaCursos []curso

var lSemestre listaCursos

func (l *listaCursos) agregarCurso(nombre string, dia int, hora bool, aula string) {
	if l.buscarChoque(dia, hora, aula) == 1 {
		println("Choca horario")
	} else {
		lSemestre = append(lSemestre, curso{nombre: nombre, dia: dia, hora: hora, aula: aula})
	}
}

func (f *listaCursos) buscarChoque(dia int, hora bool, aula string) int32 { //el retorno es el índice del producto encontrado y -1 si no existe
	lista := map2(*f, func(c curso) int32 {
		if c.dia == dia {
			if c.hora == hora {
				if c.aula == aula {
					return 1
				}
			}
		}
		return 0

	})
	return reduce(lista)
}

func map2[P1, P2 any](list []P1, f func(P1) P2) []P2 {
	mapped := make([]P2, len(list))

	for i, e := range list {
		mapped[i] = f(e)
	}
	return mapped
}

func reduce(list []int32) int32 {
	var result int32 = 0
	for _, v := range list {
		result += v
	}
	return result
}
func quemados() {
	lSemestre.agregarCurso("Lenguajes de Programacion", 3, true, "Lab2")
	lSemestre.agregarCurso("Lenguajes de Programacion", 3, true, "Lab2")
	lSemestre.agregarCurso("Bases de datos II", 4, true, "Lab1")
	lSemestre.agregarCurso("Requerimientos de Software", 3, false, "A1")
	lSemestre.agregarEstdiante("Lenguajes de Programacion", "Olman", 90)
	lSemestre.agregarEstdiante("Lenguajes de Programacion", "Allan", 85)
	lSemestre.agregarEstdiante("Lenguajes de Programacion", "Steven", 90)
	lSemestre.agregarEstdiante("Lenguajes de Programacion", "Ignacio", 80)
	lSemestre.agregarEstdiante("Lenguajes de Programacion", "Olman", 90)
	lSemestre.agregarEstdiante("Lenguajes de Programacion", "Anna", 50)

	lSemestre.agregarEstdiante("Bases de datos II", "Olman", 90)
	lSemestre.agregarEstdiante("Bases de datos II", "Allan", 85)
	lSemestre.agregarEstdiante("Bases de datos II", "Steven", 90)
	lSemestre.agregarEstdiante("Bases de datos II", "Ignacio", 80)
	lSemestre.agregarEstdiante("Bases de datos II", "Olman", 90)
	lSemestre.agregarEstdiante("Bases de datos II", "Melany", 50)
	lSemestre.agregarEstdiante("Requerimientos de Software", "Olman", 90)
	lSemestre.agregarEstdiante("Requerimientos de Software", "Allan", 85)
	lSemestre.agregarEstdiante("Requerimientos de Software", "Steven", 90)
	lSemestre.agregarEstdiante("Requerimientos de Software", "Ignacio", 80)
	lSemestre.agregarEstdiante("Requerimientos de Software", "Olman", 90)
	lSemestre.agregarEstdiante("Requerimientos de Software", "Anthony", 30)

}
func (l *listaCursos) buscarCurso(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaCursos) agregarEstdiante(nombreC string, nombreE string, nota int) {

	var curs = l.buscarCurso(nombreC)
	if (curs) != -1 {
		var est = -1
		var i int
		for i = 0; i < len((*l)[curs].lEstudiantes); i++ {
			if (*l)[curs].lEstudiantes[i].nombre == nombreE {
				est = i
			}
		}
		if est != -1 {
			fmt.Println("Ya existe estudiante en curso")
		} else {
			(*l)[curs].lEstudiantes = append((*l)[curs].lEstudiantes, estudiante{nombre: nombreE, nota: nota})
		}

	} else {
		fmt.Println("Ya existe estudiante en curso")
	}
}

func (l *listaCursos) notasC(nombre string) ([]estudiante, []estudiante) {
	var cur = lSemestre.buscarCurso(nombre)
	var lE = (*l)[cur].lEstudiantes

	notaA := filter2(lE, func(e estudiante) bool {
		return e.nota >= 70
	})
	notaR := filter2(lE, func(e estudiante) bool {
		return e.nota < 70
	})
	return notaA, notaR

}

func promedioCurso(nombre string) float32 {
	var cur = lSemestre.buscarCurso(nombre)
	var lE = lSemestre[cur].lEstudiantes
	lNotas := map2(lE, func(e estudiante) int32 {
		return int32(e.nota)
	})
	return float32(reduce(lNotas)) / float32(len(lNotas))
}

func filter2[P1 any](list []P1, f func(P1) bool) []P1 {
	filtered := make([]P1, 0)

	for _, element := range list {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}
func main() {
	quemados()
	fmt.Println("Lista de cursos", lSemestre)
	notaA, notaR := lSemestre.notasC("Lenguajes de Programacion")
	fmt.Println("Aprobados", notaA)
	fmt.Println("Reprobados", notaR)
	fmt.Println("Promedio", promedioCurso("Lenguajes de Programacion"))

}
