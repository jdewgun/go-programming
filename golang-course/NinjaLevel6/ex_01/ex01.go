package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomInteger := foo()
	barString, barInt := bar()

	fmt.Println(randomInteger)
	fmt.Println(barString, barInt)
}

func foo() int {
	x := rand.Intn(100)

	return x
}

func bar() (int, string) {
	x, y := rand.Intn(100), "Value set"

	return x, y
}
