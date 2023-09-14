//Ejercicio 2 funcional

let containsWord (word: string) (cadena: string) =
    cadena.Contains(word)

//Funcion filter
let rec filterFunc (predicate: 'a -> bool) (list: 'a list) =
    match list with
    | [] -> []
    | head :: tail ->
        //Si se cumple con el criterio
        if predicate head then
            head :: filterFunc predicate tail
        //Continua filtrando
        else
            filterFunc predicate tail

//Función que a partir de una lista de cadenas de parámetro, filtra aquellas que contengan una subcadena que el usuario indique en otro argumento.
let filterStringList (word: string) (stringList: string list) =
    let isWord stringElement =
        containsWord word stringElement
    
    filterFunc isWord stringList

//Prueba
let stringList = ["la casa"; "el perro"; "pintando la cerca"]
let word = "la"
let test = filterStringList word stringList

printfn "%A" test
