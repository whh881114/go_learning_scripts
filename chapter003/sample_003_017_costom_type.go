package main

import (
	"fmt"
	"reflect"
)

type (
	字符串 string
)

func main() {
	var b 字符串
	b = "这是中文。"
	fmt.Println(reflect.TypeOf(b), b)
	a := "这也是中文。"
	fmt.Println(reflect.TypeOf(a), a)
	// fmt.Println(b + a)
	fmt.Println(string(b) + a)

	c := 10
	fmt.Printf("%X\n", &c)
}
