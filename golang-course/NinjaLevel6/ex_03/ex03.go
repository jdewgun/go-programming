package main

import (
	"fmt"
)

func main() {
	defer foo()
	fmt.Println("Main Func running")
}

func foo() {
	defer func() {
		fmt.Println("DEFER Func Ran")
	}()

	fmt.Println("Running Func foo")
}