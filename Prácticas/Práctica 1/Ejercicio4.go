//Ejercicio 4 - Casey Baeza

package main

import (
	"fmt"
)

// Shoe struct
type Calzado struct{
	brand    string
	price    float64
	size     int
	quantity int
}

// Sells shoes if available
func sell(inventory []Calzado, brand string, size int, quantity int) {
	for x, shoe := range inventory {
		//checks if shoe is available in the store
		if shoe.brand == brand && shoe.size == size{
			//checks if shoe is in stock
			if shoe.quantity >= quantity{
				fmt.Printf("Venta exitosa\n")
				//updates quantity
				inventory[x].quantity -= quantity
				return
			}else{
				fmt.Printf("Venta fallida: calzado agotado\n")
				return
			}
		}
	}
	fmt.Printf("Venta fallida: calzado no disponible\n")
}

func main(){
	//hard code
	inventory := []Calzado{
		{brand: "Nike 1", price: 60000, size: 35, quantity: 11},
		{brand: "TW", price: 33000, size: 32, quantity: 5},
		{brand: "Adidas", price: 65000, size: 40, quantity: 12},
		{brand: "Nike 2", price: 38050, size: 37, quantity: 3},
		{brand: "Champion", price: 70000, size: 37, quantity: 2},
		{brand: "Puma", price: 38050, size: 38, quantity: 3},
		{brand: "Champion 1", price: 36000, size: 36, quantity: 2},
		{brand: "Puma 1", price: 37500, size: 39, quantity: 5},
		{brand: "SR", price: 33000, size: 32, quantity: 5},
		{brand: "NY", price: 20000, size: 42, quantity: 1},
	}

	//Test
	sell(inventory, "Adidas", 40, 13)
	sell(inventory, "Nike", 42, 3)
	sell(inventory, "Puma", 38, 2)
	sell(inventory, "QA", 39, 1)
	sell(inventory, "Puma", 38, 6)
}
