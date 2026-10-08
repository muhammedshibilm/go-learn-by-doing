package main

import "fmt"

// TODO: create the four variables and return them.
func Profile() (string, int, float64, bool) {

	var name string = "Gopher"
	age := 10
	height := 1.75
	isStudent := true

	return name, age, height, isStudent
}

func main() {
	name, age, height, isStudent := Profile()
	fmt.Printf("Name: %s, Age: %d, Height: %.2f, Is Student: %t\n", name, age, height, isStudent)
}
