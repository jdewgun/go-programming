package main

import "fmt"

var x int = 120
var y string = "Hello, World"
var z bool = true


func main() {
	fmt.Printf("The type of variable x is %T and its value is %d \n", x, x)
	fmt.Printf("The type of variable y is %T and its value is %s \n", y, y)
	fmt.Printf("The type of variable z is %T and its value is %t \n", z, z)
}


