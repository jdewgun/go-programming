package main

import (
	"fmt"
)

func main() {
	var favFood string

	favFood = "pasta"

	switch favFood {
	case "pasta":
		fmt.Println("Mushroom & Cream")
	case "pizza":
		fmt.Println("Pericolosa Pizza")
	case "chinese":
		fmt.Println("Riz Cantonais")
	case "lebanese":
		fmt.Println("Hummus Bro?")
	}
}