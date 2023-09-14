module RutaCorta

let labyrinth = [
    (1, [7]); (2, [8; 3]); (3, [4; 9]); (4, [3; 10]); (5, [6; 11]); (6, [5]);
    (7, [13]); (8, [2; 9]); (9, [3; 8]); (10, [4; 16]); (11, [5; 17]);
    (12, [18]); (13, [7; 14]); (14, [13; 15; 20]); (15, [14; 21]);
    (16, [10; 22]); (17, [11; 23]); (18, [12; 24]); (19, [25]);
    (20, [14; 26]); (21, [15; 22]); (22, [16; 21]); (23, [17; 29]);
    (24, [18; 30]); (25, [19; 31]); (26, [20; 27]); (27, [26; 28]);
    (28, [27; 29; 34]); (29, [23; 28]); (30, [24; 36]); (31, [25; 32]);
    (32, [31; 33]); (33,[32; 34]);(34,[28; 33; 35]); (35,[34; 36]); (36,[30; 35])
]

//Laberinto Prueba 2
let labyrinth2 = [
    (2, [8; 3]); (3, [4; 9]); (4, [3; 10; 5]); (5, [6; 11]); (6, [5]);
    (7, [13]); (8, [2; 9; 14]); (9, [3; 8]); (10, [4; 16]); (11, [5; 17]);
    (12, [18]); (13, [7; 14]); (14, [13; 15; 20]); (15, [14; 21]);
    (16, [10; 22]); (17, [11; 23]); (18, [12; 24]); (19, [25]);
    (20, [14; 26]); (21, [15; 22]); (22, [16; 21]); (23, [17; 29]);
    (24, [18; 30]); (25, [19; 31]); (26, [20; 27; 32]); (27, [26; 28]);
    (28, [27; 29; 34]); (29, [23; 28]); (30, [24; 36]); (31, [25; 32]);
    (32, [31; 33]); (33,[32; 34]);(34,[28; 33; 35]); (35,[34; 36]); (36,[30; 35])
]



// Función para verificar si un elemento está en una lista
let miembro elem lista =
    List.exists (fun x -> x = elem) lista

// Función para generar vecinos
let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | (head, neighbors) :: rest ->
        if head = nodo then neighbors
        else vecinos nodo rest

// Función para extender una ruta
let extender ruta grafo = 
    List.filter
        (fun x -> x <> [])
        (List.map  (fun x -> if (miembro x ruta) then [] else x::ruta) (vecinos (List.head ruta) grafo))
        
// Función para obtener todas las rutas
let allPaths ini fin grafo =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (List.head (List.head ruta)) = fin then
            [List.rev (List.head ruta)] @ (prof_aux fin ((extender (List.head ruta) grafo) @ (List.tail ruta)) grafo)       
        else
            prof_aux fin ((List.tail ruta) @ (extender (List.head ruta) grafo)) grafo
    prof_aux fin [[ini]] grafo
    
// Función principal de búsqueda en profundidad (ruta mas corta)
let rec prof2 ini fin grafo =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (List.head (List.head ruta)) = fin then
            List.rev (List.head ruta)     
        else
            prof_aux fin ((List.tail ruta) @ (extender (List.head ruta) grafo)  ) grafo
    prof_aux fin [[ini]] grafo







        
    
