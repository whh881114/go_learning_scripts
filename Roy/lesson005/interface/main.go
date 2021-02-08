package main

import "fmt"

// 接口值由两部分组成，一个是类型，一个是值。

func main() {
	var x interface{} // <type, value>
	var a int64 = 100
	var b int32 = 10
	var c int8 = 1

	x = a // <int64, 100>
	x = b // <int32, 10>
	x = c // <int8, 1>
	fmt.Println(x)

	// 类型断言
	// value, ok := x.(int)
	// if ok {
	// 	fmt.Printf("x存的是一个int类型，值是%v。\n", value)
	// }

	x = 6.66
	value, ok := x.(float64)
	fmt.Printf("ok:%t, value:%v, type:%T\n", ok, value, value)
}
