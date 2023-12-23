package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}
	peter := person{firstName: "Peter", lastName: "Parker"}
	fmt.Println(alex, peter)

	// go assigns "zero values" to unassigned fields when initializing a struct
	var gerry person
	fmt.Println(gerry)
	fmt.Printf("%+v", gerry)

	gerry.firstName = "Gerry"
	gerry.lastName = "Goobler"
	fmt.Printf("%+v", gerry)
}
