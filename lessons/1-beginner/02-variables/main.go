package main

import "fmt"

// TODO: create the four variables and return them.
func Profile() (string, int, float64, bool) {
	return "", 0, 0.0, false
}

func main() {
	name, age, height, isStudent := Profile()
	fmt.Printf("Name: %s, Age: %d, Height: %.2f, Is Student: %t\n", name, age, height, isStudent)
}
