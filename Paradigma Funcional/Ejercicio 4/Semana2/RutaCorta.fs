module RutaCorta

// Grafo con pesos
let grafo = [
    ("i", [("a", 4); ("b", 2)]);
    ("a", [("i", 1); ("c", 5); ("d", 3)]);
    ("b", [("i", 1); ("c", 7); ("d", 5)]);
    ("c", [("a", 4); ("b", 4); ("x", 8)]);
    ("d", [("a", 2); ("b", 4); ("f", 2)]);
    ("f", [("d", 1)]);
    ("x", [("c", 6)])
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

// Función para extender ruta
let extender (ruta, pesoTotal) grafo =
    let nodo_actual, _ = List.head ruta
    [ for (nodo, peso) in vecinos nodo_actual grafo do
        if not (List.exists (fun (n, _) -> n = nodo) ruta) then
            yield ((nodo, peso)::ruta, pesoTotal + peso) ]

// Función principal de búsqueda en profundidad con pesos
let rec prof2 ini fin grafo =
    let rec prof_aux fin rutas grafo shortestPath =
        match rutas with
        | [] -> shortestPath
        | (ruta_actual, peso_actual) :: rest ->
            if fst (List.head ruta_actual) = fin && peso_actual < snd shortestPath then
                let shortestPath' = (ruta_actual, peso_actual)
                prof_aux fin rest grafo (if peso_actual < snd shortestPath then shortestPath' else shortestPath)
            else
                let nuevas_rutas = extender (ruta_actual, peso_actual) grafo
                prof_aux fin (rest @ nuevas_rutas) grafo shortestPath
    prof_aux fin [([(ini, 0)], 0)] grafo ([], System.Int32.MaxValue)