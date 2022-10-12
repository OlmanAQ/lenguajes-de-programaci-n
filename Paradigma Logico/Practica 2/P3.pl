/*3)	Implemente la función aplanar. Esta función recibe una lista con múltiples listas anidadas como elementos y devuelve una lista con los mismos elementos de manera lineal (sin listas). Ej: 
aplanar([1,2,[3,[4,5],[6,7]]],X).
True
X=[1,2,3,4,5,6,7].
*/

aplanar([],[]).
aplanar([X|Y],Z):-aplanar(X,A),aplanar(Y,B),append(A,B,Z).
aplanar(X,[X]).
