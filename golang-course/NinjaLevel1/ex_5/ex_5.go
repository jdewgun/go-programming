package main

import (
	"fmt"
)

type developerTea int

var x developerTea
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 3000
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}