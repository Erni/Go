package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	info      contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}

func main() {
	p := person{
		firstName: "Alex",
		lastName:  "Anderson",
		info: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 345345,
		},
	}

	p.updateName("Ernesto")

	fmt.Println(p)

	p.print()

	a := "Bill"
	fmt.Println(&a)
	fmt.Println(&p)
}
