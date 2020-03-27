package main

import (
	"fmt"
	"runtime"
)

func main() {

	c := make(chan int)

	fmt.Println("Routines Running, ", runtime.NumGoroutine())

	// Launching 10 Goroutines which add 10 values each times
	for i := 0; i < 10; i++{
		go func(){
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}

	fmt.Println("Routines Running, ", runtime.NumGoroutine())

	// Printing out values
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("Routines Running, ", runtime.NumGoroutine())
	fmt.Println("Exiting Main.")

}