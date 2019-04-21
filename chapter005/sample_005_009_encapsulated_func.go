package main

import "fmt"

func main() {
	fmt.Printf("%d\n", Add()(3))
	fmt.Printf("%d\n", Add2(6)(3))
}

// Add无参数函数，返回值是一个匿名函数，匿名函数返回一个int类型的值
func Add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

// Add2无参数函数，返回值是一个匿名函数，匿名函数返回一个int类型的值
func Add2(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
