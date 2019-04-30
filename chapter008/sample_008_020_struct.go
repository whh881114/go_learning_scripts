package main

import (
	"fmt"
	"math"
)

// 一个已知的具名结构体
type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (n *NamePoint) Abs2() float64 {
	return n.Abs() * 100
}

// NamePoint结构体内部包含匿名字段Point
type NamePoint struct {
	Point
	name string
}

func main() {
	n := &NamePoint{Point{3, 4}, "Pythongoooo"} // n是指针
	fmt.Println(n.Abs())
	fmt.Println(n.Abs2())
}
