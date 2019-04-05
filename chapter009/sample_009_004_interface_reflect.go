package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1 // interface可以接受任何值，此时把sq1赋值，然后areaIntf就可以调Area()函数。

	// areaIntf的类型是否是Square，areaIntf.(*Square)这样的判断方法是Go语言内置的一种方法，函数中要接值的类型。
	// 此例中的Square返回的是一个指针，需要使用*来做一次运算。
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("areaIntf的类型是：%T\n", t)
	}

	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("areaIntf的类型是：%T\n", u)
	} else {
		fmt.Printf("areaIntf不含类型为Circle的变量")
	}
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}
