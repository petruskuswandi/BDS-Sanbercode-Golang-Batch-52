package main

import "fmt"

func main() {
	type NoKTP string

	var ktpPetrus NoKTP = "11111111111"

	var contoh string = "22222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpPetrus)
	fmt.Println(contohKtp)
}
