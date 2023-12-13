package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	fmt.Println("Selesai")

	names := []string{"Chloe", "Queensha", "Kuswandi"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for index, value := range names {
		fmt.Println("index", index, "=", value)
	}
}
