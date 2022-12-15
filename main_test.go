package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	// Set up
	// ...

	// Run test cases
	t.Run("String contains all letters", func(t *testing.T) {
		stringVal := "abcdefghijklmnopqrstuvwxyz"
		result := containsAllLetters(stringVal)

		if !result {
			t.Errorf("expected result to be true, got false")
		}
	})

	t.Run("String does not contain all letters", func(t *testing.T) {
		stringVal := "abcdefghijklmnopqrstuvwxy"
		result := containsAllLetters(stringVal)

		if result {
			t.Errorf("expected result to be false, got true")
		}
	})

	// Tear down
	// ...

}
