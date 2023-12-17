package main

import "fmt"

var Println = fmt.Println

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	Println(sumAll(10, 10, 10))
	Println(sumAll(10, 10, 10, 10, 10, 10))
	Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10, 10))

	numbers := []int{10, 10, 10, 10}
	Println(sumAll(numbers...))
}
