package main

import "testing"

func TestHello(t *testing.T) {

	got := Hello("Doot")
	expected := "Hello, Doot"

	if got != expected {
		t.Errorf("Got %q expected %q", got, expected)

	}
}
