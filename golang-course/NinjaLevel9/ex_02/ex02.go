package main

import (
	"fmt"
)

type person struct {
	firstName string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello World")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		firstName: "James",
	}

	saySomething(&p1)

	// does not work
	// saySomething(p1)

	p1.speak()
}