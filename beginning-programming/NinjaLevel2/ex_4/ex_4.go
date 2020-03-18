package main

import (
	"fmt"
)

func main() {
	var a int
	a = 42

	fmt.Printf("Original Val: %v\n\n", a)
	fmt.Printf("Decimal Val: %d\nBinary Val: %b\nHex Val: %#X\n", a, a, a)

	fmt.Printf("\n\nBitShifting 1 Position Left\n\n")

	b := a << 1
	fmt.Printf("BitShifted Val: %v\n\n", b)

	fmt.Printf(
		"Bitshifted Decimal Val: %d\nBitshifted Binary Val: %b\nBitshifted Hex Val: %#X\n",
		b, b, b,
	)
}