package main

import "fmt"

var Println = fmt.Println

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Petrus", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
