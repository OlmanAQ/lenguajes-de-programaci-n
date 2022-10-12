%        a ---- c ---- x
%      /   \   /
%     /	    \ /
%   i        X
%     \     / \
%      \   /   \
%        b ---- d ---- f
%

conectado(i,a,5).
conectado(i,b,10).

conectado(a,i,3).
conectado(a,c,4).
conectado(a,d,7).

conectado(b,i,9).
conectado(b,c,5).
conectado(b,d,7).

conectado(c,a,3).
conectado(c,x,8).
conectado(c,b,3).

conectado(d,b,2).
conectado(d,a,1).
conectado(d,f,9).

conectado(x,c,5).
conectado(f,d,6).


ruta(Ini,Fin,Ruta,Peso):- ruta2(Ini,Fin,[],Ruta,Peso).

ruta2(Fin,Fin,_,[Fin],0).
ruta2(Ini,Fin,Visitados,[Ini|Resto],Peso):-

    conectado(Ini,Alguien,Peso2),
    not(member(Alguien,Visitados)),
    ruta2(Alguien,Fin,[Ini|Visitados],Resto,P),
    Peso is P + Peso2.

listaRutas(Ini,Fin,ListaRuta):- findall([Ruta,Peso],ruta(Ini,Fin,Ruta,Peso),ListaRuta).

rutaCorta(Ini,Fin):-
    listaRutas(Ini,Fin,ListaRuta),
    rutaCorta2(ListaRuta,[],100,RutaCorta,Peso),
    write('La ruta corta es: '),
    write(RutaCorta),
    nl,
    write('El peso de la ruta es: '),
    write(Peso).

rutaCorta2([],RF,PF,RF,PF):-!.
rutaCorta2([[R,P]|T],_,Pa,RutaCorta,Peso):-
    P=<Pa,
    rutaCorta2(T,R,P,RutaCorta,Peso).

rutaCorta2([_|T],Ra,Pa,RutaCorta,Peso):-
    rutaCorta2(T,Ra,Pa,RutaCorta,Peso).


