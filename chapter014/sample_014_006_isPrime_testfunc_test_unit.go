package main

import (
	"goTraining/GO语言编程-黄靖钧"
	"testing"
)

// TestPrime测试
func TestPrime(t *testing.T) {
	var primeTests = []struct {
		input    int  // 输入
		expected bool // 期望结果
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
	}

	for _, tt := range primeTests {
		actual := GO语言编程_黄靖钧.isPrime(tt.input)
		if actual != tt.expected {
			t.Errorf("%d %v %v", tt.input, actual, tt.expected)
		}
	}
}
