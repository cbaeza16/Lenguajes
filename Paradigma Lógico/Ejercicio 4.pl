% Defina un predicado llamado partir(Lista, Umbral, Menores, Mayores) para dividir una lista respecto un umbral dado, dejando los valores menores a la izquierda y los mayores a la derecha.

%Caso base
partir([], _, [], []).
%Si x es menor o igual que el umbral
partir([X|Resto], Umbral, [X|Menores], Mayores) :-
    X =< Umbral,
    partir(Resto, Umbral, Menores, Mayores).
%Si x es mayor al umbral
partir([X|Resto], Umbral, Menores, [X|Mayores]) :-
    X > Umbral,
    partir(Resto, Umbral, Menores, Mayores).

%Pruebas
%?- partir([2,7,4,8,9,1],6, Menores, Mayores).
