package main

// Variables can be initialized outside a function
// A value can only be assigned in a function
var deckSize int

func main() {
	cards := deck{"Ace of Diamonds", "Five of Diamonds"}
	cards = append(cards, "Six of Spades")

	cards.print()
}
