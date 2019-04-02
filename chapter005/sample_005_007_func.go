package main

import "fmt"

func main() {
	ageArr := []int{7, 9, 3, 5, 1}
	f1(ageArr...)
}

func f1(arr ...int) {
	f2(arr...)
	fmt.Println()
	f3(arr)
}

func f2(arr ...int) {
	for _, char := range arr {
		fmt.Printf("%d ", char)
	}
}

func f3(arr []int) {
	for _, char := range arr {
		fmt.Printf("%d ", char)
	}
}

// 函数f2和f3的最大区别在于调用方式:
// f2(1, 3, 7, 13) //可交参数，随便写
// f3([]int{1, 3, 7, 13}) //需妥加上[]int{}未构造一个数组切片实例
