package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
	age int
}

func changeMe(pers *person) {
	// pers.firstName = "John"
	(*pers).firstName = "BigJohn"
}

func main() {
	person1 := person{
		firstName: "Robin",
		lastName: "Hood",
		age: 30,
	}

	fmt.Println(person1)

	changeMe(&person1)

	fmt.Println(person1)
}