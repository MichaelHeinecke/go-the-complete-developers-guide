package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

// Interfaces are not generic types. Other languages have generic
// types, but Go does not. Interfaces are "implicit". We don't
// manually have to say that a type satisfies an interface. If a
// type has the required methods, it will automatically satisfy
// the interface.

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// This function can take any type that implements the getGreeting()
// function.
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Receiver types are sufficient if variables are not used in the
// function body.
func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating a spanish greeting
	return "Hola!"
}
