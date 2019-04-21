package main

import (
	"fmt"
)

type ExpStruct struct {
	Mi1 int
	Mf1 float32
}

func main() {
	struct1 := new(ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.

	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)
}
