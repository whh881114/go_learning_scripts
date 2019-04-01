package main

import (
	"fmt"
	"math"
)

var Pi float64

func init() {
	Pi = 4 * math.Atan(1) // 在init函数中计算Pi的值
}

func main() {
	DPi := Pi * Pi
	fmt.Println(Pi, DPi)
}
