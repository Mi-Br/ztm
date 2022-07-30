//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	price float32
	name  string
}

func Total(p [4]Product) float32 {
	var sum float32
	for _, v := range p {
		sum = sum + v.price
	}
	return sum
}
func main() {

	var shoppingList [4]Product = [4]Product{
		{name: "Vanila ice cream", price: 3.32},
		{name: "Beef from greef", price: 7.78},
		{name: "Sousage brawt", price: 5.64},
	}

	fmt.Println("last item on the list", shoppingList[len(shoppingList)-1])
	fmt.Println("total capacity", len(shoppingList))
	fmt.Println("total price", Total(shoppingList))
	shoppingList[3] = Product{name: "avocado", price: 9.99}
	fmt.Println("last item on the list", shoppingList[len(shoppingList)-1])
	fmt.Println("total capacity", len(shoppingList))
	fmt.Println("total price", Total(shoppingList))
}
