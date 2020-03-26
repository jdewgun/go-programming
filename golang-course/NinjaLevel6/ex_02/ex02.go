package main

import (
	"fmt"
)

func main() {
	testArray := []int{5, 4, 3, 2, 1}

	arraySum1 := foo(testArray...)
	arraySum2 := bar(testArray)

	fmt.Println("Return value from Function foo:", arraySum1)
	fmt.Println("Return value from Function bar:", arraySum2)


}

func foo(x ...int) int {
	var total int

	for _, val := range x {
		total += val
	}

	return total

}

func bar(x []int) int {

	var sum int

	for _, val := range x {
		sum += val
	}

	return sum
}