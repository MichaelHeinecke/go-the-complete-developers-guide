package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// equivalent to above code
type person2 struct {
	firstName   string
	lastName    string
	contactInfo // declares both, field name and type
}

func main() {
	alex := person{"Alex", "Anderson", contactInfo{}}
	peter := person{firstName: "Peter", lastName: "Parker"}
	fmt.Println(alex, peter)

	// go assigns "zero values" to unassigned fields when initializing a struct
	var gerry person
	fmt.Println(gerry)
	fmt.Printf("%+v", gerry)

	gerry.firstName = "Gerry"
	gerry.lastName = "Goobler"
	fmt.Printf("%+v", gerry)

	jim := person{
		firstName: "Jim",
		lastName:  "Jones",
		contact: contactInfo{
			email:   "jim@abc.com",
			zipCode: 12345,
		},
	}
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
