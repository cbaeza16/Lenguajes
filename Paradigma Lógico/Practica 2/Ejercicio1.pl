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
menor([[Ruta1,Peso1]|Resto],[Ruta2,Peso2]):-
    menor(Resto,[Ruta2,Peso2]),
    Peso1>=Peso2.
menor([[Ruta1,Peso1]|Resto],[Ruta1,Peso1]):-
    menor(Resto,[_,Peso2]),
    Peso1<Peso2.

%Consulta: ruta_corta(i,f,Ruta,Peso).
