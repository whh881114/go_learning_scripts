package main

import "fmt"

func main() {
	// 手动填写这些参数
	age := ageMinOrMax("min", 1, 3, 2, 0)
	fmt.Printf("年龄最小的是%d岁。\n", age)

	// 数组作为参数
	ageArr := []int{7, 9, 3, 5, 1}
	age = ageMinOrMax("max", ageArr...)
	fmt.Printf("年龄最大的是%d岁。\n", age)
}

func ageMinOrMax(m string, a ...int) int {
	if len(a) == 0 {
		return 0
	}

	if m == "max" {
		max := a[0]
		for _, v := range a {
			if v > max {
				max = v
			}
		}
		return max
	} else if m == "min" {
		min := a[0]
		for _, v := range a {
			if v < min {
				min = v
			}
		}
		return min
	} else {
		e := -1
		return e
	}
}
