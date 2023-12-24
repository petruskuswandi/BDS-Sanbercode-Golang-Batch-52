package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	petrus := Man{"Petrus"}
	petrus.Married()

	fmt.Println(petrus.Name)
}
