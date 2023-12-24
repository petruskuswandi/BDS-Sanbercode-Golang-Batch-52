package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(err bool) {
	defer endApp()

	if err {
		panic("Ups Error")
	}

	// cara salah karena panic langsung berhenti dan tidak menjalankan program dibawahnya
	// message := recover()
	// fmt.Println("terjadi panic", message)
}

func main() {
	runApp(true)
	fmt.Println("Petrus Kuswandi")
}
