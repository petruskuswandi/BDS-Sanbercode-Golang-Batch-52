package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Chloe"
	middleName = "Queensha"
	lastName = "Kuswandi"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
