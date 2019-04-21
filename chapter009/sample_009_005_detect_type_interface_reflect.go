package main

import (
	"fmt"
	"math"
	"reflect"
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

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	var areaIntf Shaper
	dataType := new(Square)
	dataType.side = 5
	areaIntf = dataType

	t := areaIntf.(*Square)
	fmt.Printf("Square类型的%T值为：%v\n", t, t)

	// 在type-switch中不允许有fallthrough，在这个测试中，case是的内容要是string，所以在这里用到fmt.Sprintf功能，将其转换为string。
	tString := fmt.Sprintf("%T", t)
	fmt.Println(reflect.TypeOf(tString), tString)

	switch tString {
	case "*main.Square":
		fmt.Printf("Square类型的%T值为：%v\n", t, t)
	case "*main.Circle":
		fmt.Printf("Circle类型的%T值为：%v\n", t, t)
	case "nil":
		fmt.Printf("nil值：发生了意外。\n")
	default:
		fmt.Printf("未知类型 %T。\n", t)
	}
}
