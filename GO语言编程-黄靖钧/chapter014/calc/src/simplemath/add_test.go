package simplemath

import (
	"goTraining/GO语言编程-黄靖钧"
	"testing"
)

func TestAdd(t *testing.T) {
	r := GO语言编程_黄靖钧.Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	}
}
