%Defina la relación subconj(S, S1), donde S y S1 son listas representando conjuntos, que es verdadera si S1 es subconjunto de S.

%Caso base
subconj(_, []).

%Se comprueba si el primer elemento de S1 está en S recursivamente
subconj(S, [X|S1]) :-
    member(X, S),
    subconj(S, S1).

%Pruebas
%?- subconj([0, 6, 3, 4], [1, 3]).
%?- subconj([1, 5, 7, 4], [5, 4]).
