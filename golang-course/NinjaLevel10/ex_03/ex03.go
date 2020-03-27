package main

import (
	"fmt"
)

func main() {
	c := gen()
	recieve(c)

	fmt.Println("Exiting Main.")
}

// Func which recieves all the values from the channel
func recieve(c <-chan int) {
	for val := range c {
		fmt.Println(val)
	}
}

// Func which generates the channel
func gen() <-chan int {
	c := make(chan int)

	go func(){
		for i := 0; i < 100; i++ {
		c <- i
		}
		close(c)
	}()

	return c
}
