package main

import (
	"fmt"
)

func main() {

	slice1 := []string{"James", "Bond", "Shaken, not stirred"}
	slice2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(slice1)
	fmt.Println(slice2)

	sliceRecursive := [][]string{slice1, slice2}
	fmt.Println(sliceRecursive)


	for i, sliceObject := range sliceRecursive {
		fmt.Println("Record: ", i)
		for j, val := range sliceObject {
			fmt.Printf("\t Index Position: %v \t Value on Position: \t %v \n", j, val)
		}
	}
}