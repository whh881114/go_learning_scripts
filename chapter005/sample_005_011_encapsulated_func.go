package main

import "fmt"

func main() {
	var f = adder()
	// 相当于fmt.Println(adder()(1))
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}

func adder() func(int) int {
	var x int //闭包中的变量木中以在闭包函数体内声明，也可以丰外部函数声明
	return func(d int) int {
		fmt.Println("x = ", x, "d = ", d)
		x += d
		return x
	}
}

/*
三次调 用函数 adder()中变量 d 的值分别为: 1、 2 和 3。
在多次调用中，变量x的值是被保留的， 即0+1=1，然后1+2=3，最后3+3=6，
闭包函数保存并积累其中的变量的值，不管外部函数是否退出，它都能够继续操 作外部函数中的局部变量，这些局部变量同样可以是参数。
*/
