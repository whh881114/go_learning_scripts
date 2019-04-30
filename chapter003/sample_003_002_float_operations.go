package main

import "fmt"

func main() {
	var value1 float64
	value1 = 1
	value2 := 1.0000000000000000001

	if value1 == value2 {
		fmt.Println("相等")
	}
}

// 返回：相等。因为浮点精度不同会导致比较结果与预期不一样。
