//Ejercicio 3 - Casey Baeza

package main

import (
	"fmt"
)

//Rotate to left position n times
func rotateLeft(sequence* []string, n int){
	length := len(*sequence)
	//In case n is greater than length
	n = n % length

	//Temporary array
	temp := make([]string, n)
	copy(temp, (*sequence)[:n])
	copy((*sequence)[:length-n], (*sequence)[n:])
	copy((*sequence)[length-n:], temp)
}

//Rotate to right position n times
func rotateRight(sequence* []string, n int){
	length := len(*sequence)
	//In case n is greater than length

	//Temporary array
	temp := make([]string, n)
	copy(temp, (*sequence)[length-n:])
	copy((*sequence)[n:], (*sequence)[:length-n])
	copy((*sequence)[:n], temp)
}

func main(){
	sequence := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println("Secuencia Original:", sequence)

	rotation := []struct{
		direction int
		numberRot int
	}{	{0, 3},
		{1, 2},
	}

	for _, rotation := range rotation{
		newSequence := make([]string, len(sequence))
		copy(newSequence, sequence)
		
		//Left
		if rotation.direction == 0 {
			rotateLeft(&newSequence, rotation.numberRot)
		//Right
		}else if rotation.direction == 1{
			rotateRight(&newSequence, rotation.numberRot)
		}

		direction := "Izquierda"
		if rotation.direction == 1 {
			direction = "Derecha"
		}

		fmt.Printf(
			"Cantidad de rotaciones: %d, dirección: %s \nSecuencia rotada: %v\n",
			rotation.numberRot, direction, newSequence,
		)
	}
}
