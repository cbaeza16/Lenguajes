%Defina la función aplanar(L,L2). Esta función recibe una lista con múltiples listas anidadas como elementos otra lista que contendría los mismos elementos de manera lineal (sin listas)
%Caso base
aplanar([], []).
%Si elemento es lista
aplanar([List|Rest], L2) :-
    is_list(List),
    flatten(List, FlattenedList),
    flatten(Rest, FlattenedRest),
    append(FlattenedList, FlattenedRest, L2).
%Si elemento no es lista
aplanar([Element|Rest], [Element|FlattenedRest]) :-
    \+ is_list(Element),
    flatten(Rest, FlattenedRest).

%Pruebas
%?- aplanar([1,2,[3,4,5],[6,[6,7,8]],L2).
