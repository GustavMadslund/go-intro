package main

import "fmt"

func main() {

	// && ||
	rightTime := true
	rightPlace := true
	
	if rightTime && rightPlace {
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be Patient...")
	}
	
	enoughRobbers := false
	enoughBags := true
	
	if enoughRobbers || enoughBags {
		fmt.Println("Grab everything!")
	} else {
		fmt.Println("Grab whatever you can!")
	}
	
	// !
	readyToGo := true
	
	if !readyToGo {
		fmt.Println("Start the car!")
	} else {
		fmt.Println("What are we waiting for??")
	}
}