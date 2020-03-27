// Package dog allows us to more fully understand dogs.
package dog

import (
	"fmt"
)

// Years converts human years to dog years.
func Years(humanYears int) int {
	dogYears := humanYears * 7
	fmt.Println("Conversion to Dog Years complete, value is: ", dogYears)

	return dogYears
}