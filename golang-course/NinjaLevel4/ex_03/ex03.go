package main

import (
	"fmt"
)

func main() {
	arrayStuff := []int{1, 12, 23, 34, 45, 56, 67, 78, 89, 90}

	fmt.Println(arrayStuff[:5])
	fmt.Println(arrayStuff[5:])
	fmt.Println(arrayStuff[2:7])
	fmt.Println(arrayStuff[1:6])
	fmt.Println(arrayStuff)
}