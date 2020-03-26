package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin CPU", runtime.NumCPU())
	fmt.Println("Begin GoRoutine", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("Hello from Func One")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from Func Two")
		wg.Done()
	}()

	fmt.Println("Mid CPU", runtime.NumCPU())
	fmt.Println("Mid GoRoutine", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Exiting Main.")
	fmt.Println("End CPU", runtime.NumCPU())
	fmt.Println("End GoRoutine", runtime.NumGoroutine())
}