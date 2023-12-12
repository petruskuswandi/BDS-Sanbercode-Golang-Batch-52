package main

import "fmt"

func main() {
	// Cara menentukan berapa banyak data
	var names [3]string

	names[0] = "Chloe"
	names[1] = "Queensha"
	names[2] = "Kuswandi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Cara inisiasi langsung
	// var values = [3]int{90, 80, 95}

	// Cara 2 inisiasi langsung
	// var values = [3]int{
	// 	90,
	// 	80,
	// 	95,
	// }

	// fmt.Println(values)

	// Cara tidak menentukan berapa banyak array
	var values = [...]int{
		90,
		80,
		95,
		100,
		110,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}
