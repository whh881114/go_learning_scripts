package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [5]int  // 定义一个长度为5存放int类型的数组
	var b [10]int // 定义一个长度为10存放int类型的数组

	a = [5]int{1, 2, 3, 4, 5}
	b = [10]int{1, 2, 3, 4}

	fmt.Println(a)
	fmt.Println(b)

	// 声明时并初始化
	var c = [3]string{"bj", "sh", "sz"}
	fmt.Println(c)

	var d = [...]int{1, 23, 3, 4, 5, 9, 100}
	fmt.Println(d)
	fmt.Println(len(d))

	var e [20]int
	e = [20]int{1: 2, 3: 100, 10: 1000, 19: 1} // 针对具体索引进行赋值
	fmt.Println(e)
	fmt.Println(len(e))

	// 最原始
	for i := 0; i < len(e); i++ {
		fmt.Println(e[i])
	}

	for index, value := range e {
		fmt.Printf("INDEX: %03d, VALUE: %v\n", index, value)
	}

	// 作业题一：求数组所有元素的和。[1, 2, 3, ..., 100]
	// 作业题二：找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 9 ,10, 2]中找出和为8的两个元素的下标。
	//			此题中的下标分别为(0,3)，(1,2)。
	exp1 := [10]int{1, 2, 43, 59, 101, 1010, 2435, 853}
	exp1Sum := 0
	for _, value := range exp1 {
		exp1Sum += value
	}
	fmt.Printf("exp1数组元素之和：%d\n", exp1Sum)

	exp2 := [10]int{4, 8, 5, 3, 0, 2, 6, 7, 1}
	flagNum := 11
	existItem := "" // 记录已经算出对应和的元素，用于去重。
	for index, value := range exp2 {
		remainder := flagNum - value

		if !strings.Contains(existItem, string(value)) || !strings.Contains(existItem, string(remainder)) {
			for remainderIndex := 0; remainderIndex < (len(exp2)); remainderIndex++ {
				if remainder == exp2[remainderIndex] {
					fmt.Printf("两个元素和为%d所对应的数据下标为%d和%d，其值分别对应%d和%d。\n",
						flagNum, index, remainderIndex, value, remainder)
					existItem += string(value) + "_" + string(remainder)
				}
			}
		}
	}

}
