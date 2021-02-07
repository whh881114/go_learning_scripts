package math_pkg

import "fmt"

func init() {
	fmt.Println("我是一个无聊的计算器！init。")
}

func Add(x, y int) int {
	return x + y
}
