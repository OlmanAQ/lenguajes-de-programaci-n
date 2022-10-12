/*sub_conjunto. Esta función recibe dos listas y debe retornar True cuando el primer argumento es subconjunto completo del segundo y #f en caso contrario. Por ejemplo: 
sub_conjunto([],[1,2,3,4,5]).
True
sub_conjunto([1,2,3],[1,2,3,4,5]).
True
*/

sub_conjunto([],_).
sub_conjunto([X|Y],Z):- member(X,Z), sub_conjunto(Y,Z).

