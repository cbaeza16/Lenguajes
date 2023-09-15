//open Recipientes
open RutaCorta

let pruebaRutaCorta () =
    let inicio = "i"
    let fin = "f"
    match prof2 inicio fin grafo with
    | (_, System.Int32.MaxValue) -> printfn "No hay camino válido"
    | (ruta, peso) -> 
        printfn "Camino más corto de %s a %s:" inicio fin
        printfn "Ruta: %A" ruta
        printfn "Peso total: %d" peso
pruebaRutaCorta ()
