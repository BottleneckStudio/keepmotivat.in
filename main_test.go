package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := helloWorld()
	want := "Hello World"

	if got != want {
		t.Errorf("expects to be: %s, but got: %s instead.", want, got)
	}
}
