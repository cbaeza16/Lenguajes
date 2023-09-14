//Ejercicio 3 funcional

//Función que obtiene el n-esimo elemento de una lista
let n_esimo n list =
    //Crear una lista de índices
    let indexList = [0 .. List.length list - 1]
    //Obtener los elementos correspondientes
    let item = List.map (fun i -> List.nth list i) indexList
    //Encontar elemento segun indice
    match List.tryFind (fun (_, idx) -> idx = n) (List.zip item indexList) with
    | Some (item, _) -> Some item
    | None -> None


//Pruebas

printfn "%A" (n_esimo 2 [1;2;3;4;5])
printfn "%A" (n_esimo 3 [1;2;3;4;5])
printfn "%A" (n_esimo 6 [1;2;3;4;5])
