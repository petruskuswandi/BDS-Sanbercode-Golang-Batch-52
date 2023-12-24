package main

import (
	"bds-sanbercode-golang-batch-52/database"
	"bds-sanbercode-golang-batch-52/helper"
	_ "bds-sanbercode-golang-batch-52/internal"
	"fmt"
)

func main() {
	result := helper.SayHello("Petrus")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)              // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Petrus")) // tidak bisa diakses
	helper.Contoh()

	fmt.Println(database.GetDatabase())
}
