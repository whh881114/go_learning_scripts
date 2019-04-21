package main

import (
	"fmt"
)

type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("和为：%d。\n", two1.AddThem())
	fmt.Printf("将它们添加到参数：%d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("和为：%d\n", two2.AddThem())
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
