package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	
	if eludedGuards := rand.Intn(100); eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}
	
	if openedVault := rand.Intn(100); openedVault >= 70 && isHeistOn {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		fmt.Println("The vault can't be opened!")
	}
	
	if isHeistOn {
		switch leftSafely := rand.Intn(5); leftSafely {
			case 0:
				isHeistOn = false
				fmt.Println("Heist failed. Your didn't get away!")
			case 1:
				isHeistOn = false
				fmt.Println("Heist failed. Turns out vault doors don't open from the inside...")
			case 2:
				isHeistOn = false
				fmt.Println("Heist failed. You were shot dead!")
			case 3:
				isHeistOn = false
				fmt.Println("Heist failed. You dropped the money when running away!")
			default:
				fmt.Println("Start the getaway car!")
		}
	}
	
	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("You stole:", amtStolen, "dollars!")
	}
	
	
	fmt.Println(isHeistOn)
}