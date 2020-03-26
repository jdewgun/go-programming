package main

import (
	"fmt"
)

func main() {

	x := 5

	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("The type of var f %T\n", f)

	fmt.Println(x)
	fmt.Printf("The type of var x %T\n", x)

	g := f
	g()
	fmt.Printf("The type of var g %T\n", g)

	fmt.Println("Done")
}