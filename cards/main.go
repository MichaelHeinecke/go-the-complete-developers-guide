package main

import "fmt"

// Variables can be initialized outside a function
// A value can only be assigned in a function
var deckSize int

func main() {
	// Type can be omitted
	var card string = "Ace of Spaces"
	fmt.Println(card)

	// Alternative way of declaring a variable
	card2 := "Ace of Hearts"
	fmt.Println(card2)

	deckSize = 52
	fmt.Println(deckSize)

	card3 := newCard()
	fmt.Println(card3)

	printState()
}

func newCard() string {
	return "Five of Diamonds"
}
