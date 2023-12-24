package main

import "fmt"

type Customer struct {
	// Name, Address string
	Name    string
	Address string
	Age     int
}

func (cust Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", cust.Name)
}

func main() {
	var petrus Customer
	fmt.Println(petrus)

	petrus.Name = "Petrus Kuswandi"
	petrus.Address = "Indonesia"
	petrus.Age = 21

	fmt.Println(petrus)
	fmt.Println(petrus.Name)
	fmt.Println(petrus.Address)
	fmt.Println(petrus.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	budi.sayHello("Agus")
	petrus.sayHello("Agus")
	joko.sayHello("Agus")
}
