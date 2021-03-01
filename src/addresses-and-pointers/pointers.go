package main

import "fmt"

func main() {
	star := "Polaris"
	
	//Pointer could also be declare implicitly: starAddress := &star
	var starAddress *string
	starAddress = &star
	
	fmt.Println("The address of star is", starAddress)
}