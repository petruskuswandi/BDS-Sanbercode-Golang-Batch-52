package main

import "fmt"

func getFullName() (string, string) {
	return "Petrus", "Kuswandi"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName) // Petrus Kuswandi

	firstName, _ := getFullName()
	fmt.Println(firstName) // Petrus
}
