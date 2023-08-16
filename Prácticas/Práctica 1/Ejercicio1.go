//Ejercicio 1 - Casey Baeza

package main

import (
	"fmt"
	"strings"
)

func main(){

	//Hard code
	text := "Hello world, this is an example,  \n line 2"
	charCount, wordCount, lineCount := infoCount(text)

	//Results
	fmt.Printf("Número de caracteres: %d\n", charCount)
	fmt.Printf("Número de palabras: %d\n", wordCount)
	fmt.Printf("Número de líneas: %d\n", lineCount)
}

func infoCount(text string) (int, int, int){
	//number of characters
	charCount := len(text)
	//number of words
	wordCount := len(strings.Fields(text))
	//number of lines + initial line
	lineCount := 1 + strings.Count(text, "\n")

	return charCount, wordCount, lineCount
}