package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func(){
		for i := 1; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	receive(c)

	fmt.Println("Exiting Main")

}

func receive(c chan int) {
	for val := range c {
		fmt.Println(val)
	}
}