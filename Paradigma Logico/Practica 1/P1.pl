/*
eliminar elemento de una lista
*/
eliminar(E,[E|T],T).
eliminar(E,[H|T],[H|R]):- eliminar(E,T,R).

