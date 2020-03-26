package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
	age int
}

func (p person) speak() {
	fmt.Printf("Hello my name is %v %v\n", p.firstName, p.lastName)
	fmt.Printf("My age is %v\n", p.age)
}


func main() {
	person1 := person{
		firstName: "John",
		lastName: "Doe",
		age: 32,
	}

	person1.speak()

}