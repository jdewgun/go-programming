package main

import (
	"fmt"
)

const (
	a = 2018 + iota
	b = 2019 + iota
	c = 2020 + iota
	d = 2021 + iota
)

func main() {
	fmt.Println(a, b, c, d)
}