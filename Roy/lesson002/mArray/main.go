package main

import (
	"fmt"
)

func main() {
	var a [3]int
	a = [3]int{1, 2, 3}

	var b [3][2]int
	b = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}

	fmt.Println(a)
	fmt.Println(b)

	// 注意事项：多维数组除了第一层其它层都不能用...
	c := [...][2]int{
		{100, 2},
		{200, 3},
	}
	fmt.Println(c)
	fmt.Println("================================")

	// 多维数组的遍历
	for i := 0; i < (len(c)); i++ {
		for j := 0; j < len(c[i]); j++ {
			fmt.Println(c[i][j])
		}
	}
	fmt.Println("================================")

	for index, _ := range c {
		for _, value := range c[index] {
			fmt.Println(value)
		}
	}
	fmt.Println("================================")

	// 数组是值类型
	d := [2]int{1, 2}
	e := d
	e[0] = 100
	fmt.Println(d)
	fmt.Println(e)
}
