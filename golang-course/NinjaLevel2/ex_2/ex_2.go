package main

import (
	"fmt"
)

func main() {
	a := (20 == 20)
	b := (15 <= 20)
	c := (400 >= 20)
	d := (20 != 20)
	e := (21 < 20)
	f := (20 > 44)

	fmt.Println(a,b,c,d,e,f,)
}
