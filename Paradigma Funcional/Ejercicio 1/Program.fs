//Ejercicio 1 Funcional

//Funcion shift
let rec shift side n list =
    let length = List.length list
    match side with
    //Shift Izquierdo
    | "izq" when n >= length -> List.replicate length 0
    | "izq" ->
        let ceros = List.replicate n 0
        let numbers = List.skip n list
        numbers@ceros
    //Shift Derecho
    | "der" when n >= length -> List.replicate length 0
    | "der" ->
        let numbers = List.take (length - n) list
        let ceros = List.replicate n 0
        ceros@numbers
    | _ -> list

//Prueba 1
let listTest1 = [1; 2; 3; 4; 5]
let test1 = shift "izq" 3 listTest1
printfn "Prueba 1: %A" test1

//Prueba 2
let listTest2 = [1; 2; 3; 4; 5]
let test2 = shift "der" 2 listTest2
printfn "Prueba 2: %A" test2

//Prueba 3
let listTest3 = [1; 2; 3; 4; 5]
let test3 = shift "izq" 6 listTest3
printfn "Prueba 3: %A" test3
