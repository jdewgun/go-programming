package main

import (
	"fmt"
)

func main() {
	array_ := []int{0, 6, 3, 4, 2, 7, 5, 8, 1}

	fmt.Println("Before Sort: ", array_)

	sorted_array_ := selection_sort(array_)

	fmt.Println("After Sort: ", sorted_array_)
}

func selection_sort(array_ []int) []int {
	for i := 0; i < len(array_); i++ {
		min := i
		for j := i + 1; j < len(array_); j++ {
			if array_[j] < array_[min] {
				min = j
			}
		}

		temp := array_[i]
		array_[i] = array_[min]
		array_[min] = temp
	}

	return array_

}
