package main

import "fmt"

func main() {
	fmt.Println("Let's first see how", "the Println() method works.")
	fmt.Println("Notice that each statement adds a newline for us.")
	fmt.Println("There's also a default space", "between the string arguments.")

	fmt.Print("Print", "is", "different")
	fmt.Print("See?")
	
	fmt.Println()
	
	//Printf and interpolation
	animal1 := "cat"
	animal2 := "dog"
	
	fmt.Printf("Are you a %v or a %v person\n", animal1, animal2)
	
	//Diffenrent verbs
	floatExample := 1.25
	fmt.Printf("Working with a %T.\n", floatExample)
	
	yearsOfExp := 3
	reqYearsExp := 15
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years.\n", yearsOfExp, reqYearsExp)
	
	stockPrice := 3.50
	fmt.Printf("Each share of Gopher feed is $%.2f!\n", stockPrice)
	
	//Sprint and Sprintln
	step1 := "Breathe in..."
	step2 := "Breathe out..."
	
	meditation := fmt.Sprintln(step1, step2)
	fmt.Println(meditation)
	
	//Sprintf
	template := "I wish I had a %v."
	pet := "puppy"
	
	var wish string
	wish = fmt.Sprintf(template, pet)
	fmt.Println(wish)
	
	//Scan
	fmt.Println("What would you like for lunch?")
	var food string
	fmt.Scan(&food)
	fmt.Printf("Sure, we can have %v for lunch.", food)
}
