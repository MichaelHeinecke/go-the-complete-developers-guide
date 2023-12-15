package main

// Variables can be initialized outside a function
// A value can only be assigned in a function

func main() {
	cards := newDeck()
	cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
