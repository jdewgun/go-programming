package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (sq square) calculateAreaOfSquare() float64 {

	areaOfSquare := math.Exp2(sq.side)

	return areaOfSquare
}

func (c circle) calculateAreaOfCircle() float64 {

	areaOfCircle := math.Pi * math.Exp2(c.radius)

	return areaOfCircle
}


func main() {
	circleVal := circle{
		radius: 3.5,
	}

	squareVal := square{
		side: 3,
	}

	circleArea := circleVal.calculateAreaOfCircle()

	squareArea := squareVal.calculateAreaOfSquare()

	fmt.Println("The Area of Circle: ", circleArea)
	fmt.Println("The Area of Square: ", squareArea)
}