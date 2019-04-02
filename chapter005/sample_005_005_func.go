package main

import "fmt"

func isOdd(v int) bool {
	if v%2 == 0 {
		return false
	} else {
		return true
	}
}

func isEven(v int) bool {
	if v%2 == 0 {
		return true
	} else {
		return false
	}
}

// 声明一个函数类型
type boolFunc func(int) bool

// 函数作为一个参数使用
func filter(slice []int, f boolFunc) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{3, 1, 4, 5, 9, 2}
	fmt.Println("slice = ", slice)

	odd := filter(slice, isOdd)
	fmt.Println("odd = ", odd)

	even := filter(slice, isEven)
	fmt.Println("even = ", even)
}
