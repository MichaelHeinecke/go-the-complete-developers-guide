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

	// go passes by value. A copy of the struct is created at
	// a new memory location rather than updating the existing one
	jim.updateName("jimmy")
	jim.print()

	// A pointer can be used to update the existing struct
	jimPointer := &jim // `&` returns a pointer to the memory address of a variable
	jimPointer.updateName("jimmy")
	jim.print()

	// If the function receiver has pointer type, it's not necessary to
	// explicitly convert the variable to a pointer.
	jim.updateNameByPointer("jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Will update a copy of p
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// The star in the func declaration is a type description, meaning the function works with a pointer to a person
func (pointerToPerson *person) updateNameByPointer(newFirstName string) {
	// The star here is an operator, meaning the value the pointer references will be manipulated
	// It returns the value at a memory address
	(*pointerToPerson).firstName = newFirstName
}

// Go has value types and reference types.
// Value types are: int, float, string, bool, structs
// When a value type is passed to a function, a copy of the value is created at
// a different memory address.
// Reference types are: slices, maps, channels, pointers, functions
// Reference types consists of two data structures. One to represent the value
// itself, and one to
// store the actual data. E.g. a slice consists of a slice, storing data on the
// length, capacity, and a pointer to the head of an array which stores the
// actual data stored in the slice.
// When a reference type is passed to a function, the variable will be copied
// but the pointer to the data structure storing the data themselves remains the
// same. If data elements are modified, the original underlying data structure
// will be modified. Hence, there's no need to create a pointer, e.g. when
// updating the elements of slice inside a function.
