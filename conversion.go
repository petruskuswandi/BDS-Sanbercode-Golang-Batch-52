package main

import "fmt"

func main() {
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Chloe"
	var c = name[0]
	var cString = string(c)

	fmt.Println(name)
	fmt.Println(c)
	fmt.Println(cString)
}
