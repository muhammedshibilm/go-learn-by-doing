package main

import "testing"

func TestGrade(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{100, "A"},
		{90, "A"},
		{89, "B"},
		{75, "B"},
		{74, "C"},
		{60, "C"},
		{59, "D"},
		{40, "D"},
		{39, "F"},
		{0, "F"},
		{-1, "Invalid"},
		{101, "Invalid"},
	}

	for _, tt := range tests {
		got := Grade(tt.score)

		if got != tt.want {
			t.Errorf("Grade(%d) returned %q, but we expected %q",
				tt.score, got, tt.want)
		}
	}
}
