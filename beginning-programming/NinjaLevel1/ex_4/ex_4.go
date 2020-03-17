package main

import "fmt"

type developerTea int

var x developerTea

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 3000
	fmt.Println(x)
}