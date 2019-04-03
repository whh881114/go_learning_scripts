package main

import (
	"fmt"
	"reflect"
)

type number struct {
	f float32
}

type nr number // 类型的别名

func main() {
	a := number{5.0}
	b := nr{5.0}

	var c = number(b)
	var d = b.f

	fmt.Println(a, b)
	fmt.Println(reflect.TypeOf(c), c) // 这个直接赋值为number结构体，并且把f的值也赋了值。
	fmt.Println(d)                    // 这个是取结构体里变量的值。
}
