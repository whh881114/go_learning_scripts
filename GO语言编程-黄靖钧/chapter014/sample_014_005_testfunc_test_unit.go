package main

import (
	"goTraining/GO语言编程-黄靖钧"
	"testing"
)

func TestFib(t *testing.T) {
	var (
		input    = -100
		expected = 0
	)
	actual := GO语言编程_黄靖钧.Age(input)
	if actual != expected {
		t.Errorf("Fib(%d) = %d, 预期为%d", input, actual, expected)
	}
}
