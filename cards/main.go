package main

import "fmt"

// Variables can be initialized outside a function
// A value can only be assigned in a function

func main() {
	cards := newDeck()
	cards.print()

	hand, remainingCards := deal(cards, 5)
	fmt.Println("Dealt hand:")
	hand.print()
	fmt.Println("Cards remaining in deck:")
	remainingCards.print()

	hand.saveToFile("my-hand")
	handFromFile := newDeckFromFile("my-hand")
	fmt.Println("Hand after writing it to disk and reading it from disk:")
	handFromFile.print()

	fmt.Println("Cards after shuffling:")
	cards.shuffle()
	cards.print()
}
