package main

import "fmt"

func main() {
	j := 5
	a := func() func() {
		i := 10
		return func() {
			fmt.Printf("i = %d j = %d\n", i, j)
		}
	}() // 这个是进行初始化，将函数返回给变量a。

	a() // i = 10 j =5
	j = 10
	a() // i = 10 j =10
}
