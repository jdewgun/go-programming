package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// Anonymous Function, Parallel Goroutine
	go func() {
		c <- 42
	}()

	fmt.Println("Value from Anonymous Function", <-c)

	// Buffered channel
	bufferedC := make(chan int, 1)

	bufferedC <- 42


	fmt.Println("Value from Buffered Channel", <-bufferedC)
}
