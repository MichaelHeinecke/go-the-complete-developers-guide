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
}
