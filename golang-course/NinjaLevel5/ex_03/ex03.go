package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	colour string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheelDrive bool
}

func main() {

	productOne := truck{
		vehicle: vehicle{
			doors: 2,
			colour: "Red",
		},
		fourWheelDrive: true,
	}

	productTwo := sedan{
		vehicle: vehicle{
			doors: 4,
			colour: "Black",
		},
		luxury: true,
	}

	fmt.Println(productOne)
	fmt.Println(productTwo)

	fmt.Println(productOne.doors)
	fmt.Println(productTwo.doors)
}