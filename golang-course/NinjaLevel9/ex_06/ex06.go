package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Golang OS:\t", runtime.GOOS)
	fmt.Println("Golang Architecture:\t", runtime.GOARCH)
}