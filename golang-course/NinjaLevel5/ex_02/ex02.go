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

	lastNameDirectory := map[string]person{
		person1.firstName: person1,
		person2.firstName: person2,
	}

	for _, val := range lastNameDirectory {
		fmt.Println(val.firstName)
		fmt.Println(val.lastName)

		for i, iceCreamFlavour := range val.iceCreamFlavours {
			fmt.Println(i, iceCreamFlavour)
		}
		fmt.Println("************")
	}
}