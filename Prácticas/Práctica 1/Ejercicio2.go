//Ejercicio 2 - Casey Baeza

package main

import (
	"fmt"
)

// Prints "n" *
func printSymbol(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

// Prints figure
func printFigure(center int) {
	//Print upper side
	for i := 1; i <= center; i += 2 {
		space := (center - i) / 2
		for j := 0; j < space; j++ {
			fmt.Print(" ")
		}
		printSymbol(i)
	}

	//Print lower side
	for i := center - 2; i >= 1; i -= 2 {
		space := (center - i) / 2
		for j := 0; j < space; j++ {
			fmt.Print(" ")
		}
		printSymbol(i)
	}
}

func main() {
	//Elements in middle line (odd and positive)
	centerSize := 7

	printFigure(centerSize)
}
