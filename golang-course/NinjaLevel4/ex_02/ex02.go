package main

import (
	"fmt"
)

func main() {
	arrayStuff := []int{1, 12, 23, 34, 45, 56, 67, 78, 89, 90}

	for i, v := range arrayStuff {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", arrayStuff)

}