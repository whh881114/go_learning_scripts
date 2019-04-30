package simplemath

import (
	"goTraining/GO语言编程-黄靖钧"
	"testing"
)

func TestSqrt(t *testing.T) {
	v := GO语言编程_黄靖钧.Sqrt(16)
	if v != 4 {
		t.Errorf("Sqrt(16) failed. Got %v, expected 4.0.", v)
	}
}
