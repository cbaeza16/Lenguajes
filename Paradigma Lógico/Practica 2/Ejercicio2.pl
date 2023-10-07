% Definición de personas
persona(diego, [1,2,5,2,9,8,4,6,3,4]).
persona(maria, [2,2,5,2,9,8,4,6,3,5]).
persona(pedro, [1,2,3,4,5,8,4,6,3,6]).
persona(ana, [1,2,5,2,9,6,7,6,3,7]).
persona(luis, [1,2,5,4,5,6,4,6,3,8]).
persona(lucia, [1,2,1,1,1,8,4,6,3,9]).
persona(carlos, [1,7,3,2,9,8,4,6,4,0]).
persona(beatriz, [7,8,9,2,9,8,4,6,4,1]).
persona(monica, [1,2,5,2,1,1,4,6,4,2]).
persona(sara, [1,2,5,8,9,0,4,6,4,3]).

similitud_entre_cromosomas(Cromosoma1, Cromosoma2, Similitud) :-
    calcular_similitud(Cromosoma1, Cromosoma2, Similitud).

calcular_similitud([], [], 0).
calcular_similitud([X|Resto1], [Y|Resto2], Similitud) :-
    calcular_similitud(Resto1, Resto2, SimilitudResto),
    Similitud is SimilitudResto + (X * Y).

persona_similar(Cromosoma, PersonaSimilar) :-
    findall(Similitud-Nombre, (
        persona(Nombre, Cromosomas),
        similitud_entre_cromosomas(Cromosoma, Cromosomas, Similitud)
    ), ListaSimilitudes),
    max_member(_-PersonaSimilar, ListaSimilitudes).

% Prueba
prueba :-
    Muestra = [1, 2, 5, 2, 9, 8, 4, 6, 3, 5],
    persona_similar(Muestra, PersonaSimilar),
    format('La persona es ~w.', [PersonaSimilar]).

%Consulta: prueba.
