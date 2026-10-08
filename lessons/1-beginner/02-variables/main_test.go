package main

import "testing"

func TestProfile(t *testing.T) {
	name, age, height, isStudent := Profile()

	if name != "Gopher" {
		t.Errorf("name is %q, but we expected %q", name, "Gopher")
	}
	if age != 10 {
		t.Errorf("age is %d, but we expected %d", age, 10)
	}
	if height != 1.75 {
		t.Errorf("height is %v, but we expected %v", height, 1.75)
	}
	if !isStudent {
		t.Errorf("isStudent is %v, but we expected true", isStudent)
	}
}
