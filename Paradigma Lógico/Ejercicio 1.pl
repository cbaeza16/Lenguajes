%Defina sumlist(L, S) que es verdadero si S es la suma de los elementos de L.

%Caso base
sumlist([], 0).

%Sumar el primer elemento de la lista a la suma de los elementos restantes.
sumlist([X|Suma], S) :-
    sumlist(Suma, Resto),
    S is Resto + X.

%Pruebas
%?- sumlist([2,2,5],9)
%?- sumlist([2,1,5],9)
