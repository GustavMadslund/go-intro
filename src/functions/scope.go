package main

import "fmt"

//Moved fmt.Println(instructions) into startGame as instructions is not accesible from main
func startGame() {
	instructions := "Press enter to start..."
	fmt.Println(instructions)
}

func main() {
	startGame()
}