package main

import (
	"fmt"
)

func main() {
	year := 1994

	for {
		if year > 2020 {
			break
		}
		fmt.Println(year)
		year++
	}
}