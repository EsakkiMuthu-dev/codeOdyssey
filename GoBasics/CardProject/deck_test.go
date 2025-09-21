package main

import "testing"

func TestLengthOfDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 20 {
		t.Errorf(" Expected 20 but got %v", len(d))
	}
}
