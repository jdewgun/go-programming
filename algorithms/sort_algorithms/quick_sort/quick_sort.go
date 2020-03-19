package main

import (
	"math/rand"
	"fmt"
)

func main() {
	unsorted_array := []int{5,7,5,8,2,4,8,5,3,5,1}
	fmt.Println("The array before sort: ", unsorted_array)

	sorted_array := quick_sort(unsorted_array)

	fmt.Println("The array after sort: ", sorted_array)
}

func quick_sort(array_ []int) []int {

	if len(array_) <= 1 {
		return array_
	}

	median := array_[rand.Intn(len(array_))]

	low_part := make([]int, 0, len(array_))
	middle_part := make([]int, 0, len(array_))
	high_part := make([]int, 0, len(array_))


	for _, item := range array_ {
		switch {
		case item < median:
			low_part = append(low_part, item)
		case item == median:
			middle_part = append(middle_part, item)
		case item > median:
			high_part = append(high_part, item)

		}
	}

	low_part = quick_sort(low_part)
	high_part = quick_sort(high_part)

	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)

	return low_part

}