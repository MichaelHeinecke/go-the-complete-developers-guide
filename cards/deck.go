package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string

// (d deck) is called a receiver
// Variables of type deck get access to this method
// By convention the receiver is a one or two letter abbreviation of the type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
