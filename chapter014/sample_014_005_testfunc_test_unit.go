package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	var (
		input    = -100
		expected = 0
	)
	actual := Age(input)
	if actual != expected {
		t.Errorf("Fib(%d) = %d, 预期为%d", input, actual, expected)
	}
}
