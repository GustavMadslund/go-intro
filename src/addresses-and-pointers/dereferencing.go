package main

import "fmt"

func main() {
	star := "Polaris"
	
	starAddress := &star
	
	//Dereference the pointer
	*starAddress = "Sirius"
	
	fmt.Println("The actual value of star is", star)
}