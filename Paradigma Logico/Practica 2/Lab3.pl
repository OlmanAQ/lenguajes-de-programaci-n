/*
Predicado que permita retornar una lista con la mejor distribución equitativa de monedas costarricenses para realizar un cambio bancario de billetes a “menudo”. Por ejemplo, si el usuario desea cambiar 10395 colones en monedas, la aplicación debe retornar cual es la mejor distribución de monedas, que para este caso debe ser: 20 de 500, 3 de 100, 1 de 50, 1 de 25, 2 de 10 y 0 de 5, entendiendo por “mejor” aquella que priorice menor Billetes de monedas a partir de las de mayor denominación.
*/
%menudo(10395, [500, 100, 50, 25, 10, 5], X ).
%X = [20, 3, 1, 1, 2, 0]

menudo(0, [], []).
menudo(Cambio, [H|T], [X|Y]) :- 
    Cambio >= H,
    X is Cambio // H,
    Cambio1 is Cambio mod H,
    menudo(Cambio1, T, Y).
menudo(Cambio, [_|T], [0|Y]) :-
    menudo(Cambio, T, Y).




/*
-	Entradas: guacamole, ensalada, consomé, tostadas caprese
-	Carne: filete de cerdo, pollo al horno, carne en salsa
-	Pescado: tilapia, salmón, trucha
-	Postre: flan, nueces con miel, naranja confitada, flan de coco

-	Una ración de guacamole aporta 200 calorías
-	Una ración de ensalada aporta 150 calorías       
-	Una ración de consomé aporta 300 calorías          
-	Una ración de tostadas caprese aporta 250 calorías
-	Una ración de filete de cerdo aporta 400 calorías  
-	Una ración de pollo al horno aporta 280 calorías      
-	Una ración de carne en salsa aporta 320 calorías
-	Una ración de tilapia aporta 160 calorías           
-	Una ración de salmón aporta 300 calorías          
-	Una ración de trucha aporta 225 calorías
-	Una ración de flan aporta 200 calorías             
-	Una ración de nueces con miel aporta 500 calorías  
-	Una ración de naranja confitada aporta 450 calorías           
-	Una ración de flan de coco aporta 375 calorías           

*/
entrada(guacamole, 200).
entrada(ensalada, 150).
entrada(consome, 300).
entrada(tostadas_caprese, 250).

principal(filete_de_cerdo, 400).
principal(pollo_al_horno, 280).
principal(carne_en_salsa, 320).
principal(tilapia, 160).
principal(salmon, 300).
principal(trucha, 225).

postre(flan, 200).
postre(nueces_con_miel, 500).
postre(naranja_confitada, 450).
postre(flan_de_coco, 375).

%comida completa está compuesta por una entrada, un plato principal y un postre.
/*
Muestra las comidas que no superen un máximo X de calorías y que no repitan ningún elemento en su composición del plato con respecto a las obtenidas con anterioridad. 
*/
%no repetir entrada, principal y postre de la comida anterior ningún elemento en su composición del plato con respecto a las obtenidas con anterioridad.
%comida(1000, X).
%X = [guacamole, filete_de_cerdo, flan, 750]
%X = [ensalada, carne_en_salsa, nueces_con_miel, 970]



comida(Max, [E, P, Po, Calorias]) :-
    entrada(E, CaloriasE),
    principal(P, CaloriasP),
    postre(Po, CaloriasPo),
    Calorias is CaloriasE + CaloriasP + CaloriasPo,
    Calorias =< Max.



%verifica que la lista de menu no estén repetido en la comida la entrada, principal y postre en la comida
menu(Max, X) :-
    findnsols(3,[E, P, Po, Calorias], comida(Max, [E, P, Po, Calorias]), X),
    %verifica que la lista de menu no estén repetido en la comida la entrada, principal y postre en la comida
    miembro([E, P, Po, Calorias], X),
    miembro([E1, P1, Po1,_], X),
    E \= E1,
    P \= P1,
    Po \= Po1.


miembro(X, [X|_]).
miembro(X, [_|Y]) :- miembro(X, Y).



/*
La distancia de Hamming entre dos cadenas es el número de posiciones en que los correspondientes elementos son distintos. Por ejemplo, la distancia de Hamming entre “roma” y “loba” es 2 (porque hay 2 posiciones en las que los elementos correspondientes son distintos: la posición 1 y la 3). Nótese en los ejemplos, que el tamaño de las cadenas puede no ser igual y por ende debe tomarse como referencia para la comparación el tamaño de palabra menor)
*/
/*
distanciaH("romano","comino",X).
X = 2 
distanciaH("romano","camino",X).
X = 3 
distanciaH("roma","comino",X).
X = 2 
distanciaH("romano","ron",X).
X = 1 
distanciaH("romano","cama",X).
X = 2
*/

%se tiene que pasar la cadena a lista
distanciaH([], [], 0).
distanciaH([H1|T1], [H2|T2], X) :-
    H1 \= H2,
    distanciaH(T1, T2, X1),
    X is X1 + 1.
distanciaH([_|T1], [_|T2], X) :-
    distanciaH(T1, T2, X).

distanciaH(C1, C2, X) :-
    string_chars(C1, L1),
    string_chars(C2, L2),
    distanciaH(L1, L2, X).
