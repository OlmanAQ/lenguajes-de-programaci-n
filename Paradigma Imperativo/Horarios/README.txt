El laboratorio se realizo en GOLANG

Lista de cursos [{Lenguajes de Programacion 3 true Lab2 [{Olman 90} {Allan 85} {Steven 90} {Ignacio 80}]} {Bases de datos II 4 true Lab1 []} {Requerimientos de Software 3 false A1 []}]
La lista se lee Nombre, dia de 1 a 5, horario true=mañana false=tarde, aula y lista de estudiantes


funcion con errores
func (l *listaCursos) aprobadosC(nombre string) int32 {
		var cur = lSemestre.buscarCurso(nombre)
		var lE = (*l)[cur].lEstudiantes

		lista := filter2(lE), func(e estudiante) bool {
			return e.nota <70
		})

}