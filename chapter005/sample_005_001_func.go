package main

import "fmt"

func main() {
	x := 1
	y := 2
	z := 3
	var maxNum int
	maxNum = max(x, y)
	fmt.Printf("max(%d, %d) = %d\n", x, y, maxNum)

	maxNum = max(x, z)
	fmt.Printf("max(%d, %d) = %d\n", x, z, maxNum)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z))
}

// 私有函数
func max(a, b int) (maxNum int) {
	if a > b {
		return a
	} else {
		return b
	}
}

// 公有函数
func Max(a, b int) (maxNum int) {
	if a > b {
		return a
	} else {
		return b
	}
}
