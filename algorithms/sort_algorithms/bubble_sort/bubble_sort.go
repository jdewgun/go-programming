// Basic Bubble Sort Example

package main

import (
	"fmt"
)

func main() {
	an_array := []int{2,4,3,1,5}

	fmt.Println("Here is the array before the sort: ", an_array)

	sorted_array :=	bubble_sort(an_array)

	fmt.Println("Here is the array after the sort: ", sorted_array)

}

func bubble_sort(array_ []int) []int { // Bubble Sort Function

	is_swapped := true

	for is_swapped {
		is_swapped = false
		for i:=0; i < len(array_)-1; i++ {
			if array_[i + 1] < array_[i] {
				swap(array_, i, i+1)
				is_swapped = true
			}
		}
	}

	return array_
}


func swap(array_ []int, i, j int) {
	temp := array_[j]
	array_[j] = array_[i]
	array_[i] = temp
}

