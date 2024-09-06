package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		// Use %v to format the integer value of len(d). %v is the formatting directive for integers.
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
