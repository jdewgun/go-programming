package main

import (
	"fmt"
)

const (
	a = 42 // UNTYPED
	b int = 24 // TYPED
)

func main() {
	fmt.Printf("Type Of Untyped Constant: %v  %T\n", a, a)
	fmt.Printf("Type Of Typed Constant: %v  %T\n", b, b)
}