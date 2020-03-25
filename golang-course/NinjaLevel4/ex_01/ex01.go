package main

import (
	"fmt"
)

func main() {
	arrayStuff := [5]int{16, 25, 34, 43, 52}

	for i, v := range arrayStuff {
		fmt.Println(i, v)
	}

	fmt.Printf("The type of the variable is %T\n", arrayStuff)
}