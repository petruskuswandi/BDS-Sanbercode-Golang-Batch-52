package main

import (
	"fmt"
)

type Blacklist func(string) bool

func registerUSer(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUSer("petrus", blacklist)

	registerUSer("anjing", func(s string) bool {
		return s == "anjing"
	})
}
