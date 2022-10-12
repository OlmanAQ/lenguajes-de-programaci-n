/*Implementar un predicado que, a partir de una lista de cadenas de parámetro, filtre aquellas que contengan una subcadena que el usuario indique en otro argumento. Ej

sub_cadenas(“la”, [“la casa, “el perro”, “pintando la cerca”],Filtradas).
True
Filtradas = [“la casa, “pintando la cerca”]
*/

sub_cadenas(_,[],[]).
sub_cadenas(Sub,[Cadena|Resto],[Cadena|Filtradas]):-
    sub_string(Cadena,_,_,_,Sub),
    sub_cadenas(Sub,Resto,Filtradas).
sub_cadenas(Sub,[_|Resto],Filtradas):-
    sub_cadenas(Sub,Resto,Filtradas).



