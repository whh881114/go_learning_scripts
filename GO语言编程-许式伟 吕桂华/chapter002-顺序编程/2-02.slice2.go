package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)

	fmt.Println("### 查看mySlice的len和cap值：")
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	fmt.Println()

	fmt.Println("### 数组切片扩展内容，append添加一些元素：")
	mySlice = append(mySlice, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(mySlice)
	fmt.Println()

	fmt.Println("### 数组切片扩展内容，append添加另一个切片：")
	mySlice2 := []int{8, 9, 10}
	mySlice = append(mySlice, mySlice2...)
	fmt.Println(mySlice)
	fmt.Println()

	fmt.Println("### 再次查看mySlice的len和cap值：")
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
}
