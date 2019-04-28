package main

import "fmt"

func main() {
	var j int = 5

	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	// 匿名函数返回的是一个匿名函数，所以这里赋值后使用了()，就表示执行了第一层匿名函数，再次调用时就需要在后加上()再执行返回的
	// 匿名函数。

	a()
	j *= 2
	a()

	// 再换一种方式来写，先不执行第一层的匿名函数。
	b := func() func() {
		var i int = 10 // i变量是内部的，外部是无法访问此变量，保证了i的安全性。
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}

	b()()
	j *= 2
	b()()
}
