package main

import "fmt"

func main() {
	var stationName string
	var nextTrainTime int8
	var isUptownTrain bool
	
	stationName = "Union Square"
	nextTrainTime = 12
	isUptownTrain = false
	
	fmt.Println("Current station:", stationName)
	fmt.Println("Next train:", nextTrainTime, "minutes")
	fmt.Println("Is uptown:", isUptownTrain)
	
	stationName = "Grand Central"
	nextTrainTime = 3
	isUptownTrain = true
	
	fmt.Println("")
	fmt.Println("Current station:", stationName)
	fmt.Println("Next train:", nextTrainTime, "minutes")
	fmt.Println("Is uptown:", isUptownTrain)
	
	//Literals
	fmt.Println(2235 * 1231)
	
	const earthsGravity = 9.80665
	
	//Numeric types
	var jellybeanCounter int8
	fmt.Println(jellybeanCounter)
	
	var numOfFlavors int
	numOfFlavors = 57
	fmt.Println(numOfFlavors)
	
	var flavorScale float32 = 5.8
	fmt.Println(flavorScale)
	
}