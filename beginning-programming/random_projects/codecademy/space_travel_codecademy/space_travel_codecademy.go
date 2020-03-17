package main

import "fmt"


func fuelGauge(fuel int) {
  fmt.Printf("Current Fuel Level: %d", fuel)
}


func calculateFuel(planet string) int {
  var fuel int

  switch planet {
		case "Venus":
			fuel = 300000
	  	case "Mercury":
			fuel = 500000
	  	case "Mars":
			fuel = 700000
	  	default:
			fuel = 0
  }

  return fuel

}

func greetPlanet(planet string) {
  fmt.Printf("Ladies and Gentle, please welcome to %s", planet)
}

func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int

  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)

  if fuelRemaining >= fuelCost {
    greetPlanet(planet"\n")
    fuelRemaining = fuelRemaining - fuelCost
  } else {
    cantFly()
  }

  return fuelRemaining


}

func main() {

  fuel := 1000000
  planetChoice := "Venus"

  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)

}