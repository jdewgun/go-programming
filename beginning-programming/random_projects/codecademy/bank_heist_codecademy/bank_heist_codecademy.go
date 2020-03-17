/*
Codecademy Project
------------------

Project to try out some Conditional Loops

*/
package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main () {
	// Set up random seed

	rand.Seed(time.Now().UnixNano())

	// Set up base Boolean

	var isHeistOn bool

	isHeistOn = true

	// Did we get past the guards ?

	if eludedGuards := rand.Intn(100); eludedGuards >=  50 {
		fmt.Println("Looks like you've managed to make it past the guards.")
		fmt.Println("Good job, but remember, this is the first step.")
	} else {
		fmt.Println("Plan a better disguise next time?")
	}

	// Did we open the Vault ?

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("Vault is locked.")
	}

	// Did we get the money and leave safely ?

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
			case 0:
				isHeistOn = false
				fmt.Println("Caught By Police.")
			case 1:
				isHeistOn = false
				fmt.Println("Flat Tire, Caught By Police.")
			case 2:
				isHeistOn = true
				// fmt.Println("Avegned By Batman.")
				fmt.Println("We got that MONNNEEEEEEEYYY.")
			case 3:
				isHeistOn = true
				// fmt.Println("Obi-Wan Kenobi got you.")
				fmt.Println("Caught By Police.")
			case 4:
				isHeistOn = true
				fmt.Println("We got that MONNNEEEEEEEYYY.")
			}
		}

	// So how much Money did we make ?
	if amtStolen := 10000 + rand.Intn(1000000); isHeistOn {
		fmt.Printf("The Money We got is %d€", amtStolen)
	}


}