package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := "Hello World"
	want := helloWorld()

	if got != want {
		t.Errorf("expects to be: %s, but got: %s instead.", want, got)
	}
}
