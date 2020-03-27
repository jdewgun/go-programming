package dog

import (
	"fmt"
)

func Years(humanYears int) int {
	dogYears := humanYears * 7
	fmt.Println("Conversion to Dog Years complete, value is: ", dogYears)

	return dogYears
}