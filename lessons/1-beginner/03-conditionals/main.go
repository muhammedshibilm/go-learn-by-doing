package main

import "fmt"

// TODO: return the correct grade for the score.
//
// Rules:
// 90-100 -> A
// 75-89  -> B
// 60-74  -> C
// 40-59  -> D
// 0-39   -> F
// Less than 0 or greater than 100 -> Invalid
func Grade(score int) string {
	var grade string

	if score < 0 || score > 100 {
		return "Invalid"
	} else if score >= 90 {
		grade = "A"
	} else if score >= 75 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else if score >= 40 {
		grade = "D"
	} else {
		grade = "F"
	}

	return grade
}

func main() {
	fmt.Println("90:", Grade(90))
	fmt.Println("76:", Grade(76))
	fmt.Println("65:", Grade(65))
	fmt.Println("45:", Grade(45))
	fmt.Println("20:", Grade(20))
	fmt.Println("110:", Grade(110))
}
