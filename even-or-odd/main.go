package main

import "fmt"

// Assignment 1 - Even or Odd
// Create slice of ints from 0 to 10 in main function
// Iterate over slice and print out "even" for even numbers and "odd" for odd
// numbers

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	for _, v := range a {
		if v%2 == 0 {
			fmt.Printf("%v is even.\n", v)
		} else {
			fmt.Printf("%v is odd.\n", v)
		}
	}
}
