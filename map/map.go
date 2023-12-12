package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Petrus"
	// person["address"] = "Cirebon"

	person := map[string]string{
		"name":    "Petrus",
		"address": "Cirebon",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Petrus"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
