package main

import "fmt"

func main() {
	// Cara 1
	// var name string
	// name = "Petrus Kuswandi"

	// Cara 2
	// var name = "Petrus Kuswandi"

	// Cara 3
	// inisiasi ini hanya dilakukan diawal saja menggunakan := seterusnya jika ingin menginisiasi value baru maka gunakan =
	name := "Petrus Kuswandi"
	fmt.Println(name)

	name = "Chloe Queensha Kuswandi"
	fmt.Println(name)

	// Cara 4
	var (
		firstName  = "Chloe"
		middleName = "Queensha"
		lastName   = "Kuswandi"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
