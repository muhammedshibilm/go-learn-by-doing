package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()

	want := "Hello, World!"

	if got != want {
		t.Errorf("Hello() returned %q, but we expected  %q. Check your return value", got, want)
	}
}
