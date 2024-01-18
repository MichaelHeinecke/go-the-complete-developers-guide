package main

import "fmt"

func main() {
	// Ways of creating a map
	// With key value pairs
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}
	fmt.Println(colours)
	printMap(colours)

	// Empty map
	var colours2 map[string]string
	fmt.Println(colours2)

	// Make empty map
	colours3 := make(map[string]string)
	colours3["white"] = "#ffffff"
	fmt.Println(colours3)

	delete(colours3, "white")
	fmt.Println(colours3)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
