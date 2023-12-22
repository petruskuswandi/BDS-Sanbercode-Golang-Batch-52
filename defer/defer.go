package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	fmt.Println("Run application")

	// error
	logging()
}

func main() {
	runApplication()
}
