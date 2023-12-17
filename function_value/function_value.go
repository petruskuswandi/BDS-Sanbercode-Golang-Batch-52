package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	contoh := getGoodBye
	misal := getGoodBye

	Println := fmt.Println

	Println(contoh("Petrus"))
	Println(misal("Chloe"))
}
