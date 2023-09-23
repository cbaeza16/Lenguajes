%Implemente un predicado que, a partir de una lista de cadenas string,
% filtre aquellas que contengan una subcadena que el usuario indique en
% otro argumento.

% Caso base
sub_cadenas(_, [], []).

%Incluye subcadena
sub_cadenas(SubString, [String|Rest], Filtered) :-
    sub_atom(String, _, _, _, SubString),
    sub_cadenas(SubString, Rest, RestFiltered),
    Filtered = [String|RestFiltered].

%No incluye subcadena
sub_cadenas(SubString, [_|Rest], Filtered) :-
    sub_cadenas(SubString, Rest, Filtered).

%Pruebas
%?- sub_cadenas("la",["la casa", "el perro", "pintando la cerca"], Filtered).
