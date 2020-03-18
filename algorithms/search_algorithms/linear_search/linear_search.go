package main

import (
	"fmt"
)

func main() {
	array_ := []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	fmt.Println("Linear Search to be implemented on: ", array_)
	fmt.Println("Searched Item: 21")
	index := linear_search(array_, 21)

	if index == -1 {
		fmt.Println("Searched Number not in Array")
	} else {
		fmt.Println("The number is found at position: ", index)
		fmt.Println("array[", index, "] =", array_[index])
	}
}

func linear_search(array_ []int, needed_item int) int {
	for i, item := range array_ {
		if item == needed_item {
			return i
		}
	}
	return -1
}