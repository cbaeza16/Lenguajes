//open Recipientes
open RutaCorta

//Rutas posibles en el laberinto 2
let paths = allPaths 2 32 labyrinth2
printfn "--Rutas del Laberinto--\n"
List.iter (fun path -> printfn "%A" path) paths

//Ruta mas corta del laberinto 2
let shortestPath = prof2 2 32 labyrinth2
printfn "\n--Ruta Corta--\n"
printfn "%A" shortestPath