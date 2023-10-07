conectado(i,a,50).
conectado(i,b,20).
conectado(a,b,15).
conectado(a,c,25).
conectado(b,f,35).
conectado(b,c,40).
conectado(c,f,5).


ruta(Ini,Fin,Ruta, PesoTotal):- ruta2(Ini,Fin,[],Ruta, 0, PesoTotal).

ruta2(Fin,Fin,_,[Fin], PesoAcumulado, PesoAcumulado).

ruta2(Ini,Fin,Visitados,[Ini|Resto], PesoAcumulado, PesoTotal):-
    conectado(Ini,Alguien,P),
    not(member(Alguien,Visitados)),
    NuevoPesoAcumulado is PesoAcumulado + P,
    ruta2(Alguien,Fin,[Ini|Visitados],Resto, NuevoPesoAcumulado, PesoTotal).

ruta_corta(Ini, Fin,Ruta,Peso):-
    findall([R,P],ruta(Ini,Fin,R,P),Rutas),
    menor(Rutas,[Ruta,Peso]).
menor([[Ruta, Peso]],[Ruta,Peso]).
menor([[R1,P1]|Resto],[R2,P2]):-
    menor(Resto,[R2,P2]),
    P1>=P2.
menor([[R1,P1]|Resto],[R1,P1]):-
    menor(Resto,[_,P2]),
    P1<P2.

%Consulta: ruta_corta(i,f,Ruta,Peso).
