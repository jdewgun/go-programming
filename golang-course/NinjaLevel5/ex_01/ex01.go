package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
	iceCreamFlavours []string
}

func main() {
	person1 := person{
		firstName: `John`,
		lastName: `Doe`,
		iceCreamFlavours: []string{`Chocolate`,`Strawberry`,`Vanilla`},
	}

	person2 := person{
		firstName: `Person`,
		lastName: `Doe`,
		iceCreamFlavours: []string{`Vanilla`},
	}

	fmt.Printf(
		"First Name: %v \nLast Name: %v \n",
		person1.firstName,
		person1.lastName,
	)
	for i, value := range person1.iceCreamFlavours {
		fmt.Printf("Favourite Ice Cream #%v: %v\n", (i+1), value)
	}

	fmt.Printf("\n\n")

	fmt.Printf(
		"First Name: %v \nLast Name: %v \n",
		person2.firstName,
		person2.lastName,
	)
	for i, value := range person2.iceCreamFlavours {
		fmt.Printf("Favourite Ice Cream #%d: %v\n", (i+1), value)
	}


}