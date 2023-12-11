package main

import "fmt"

var Println = fmt.Println

func main() {
	Println("Benar = ", true)
	Println("Salah = ", false)
}
