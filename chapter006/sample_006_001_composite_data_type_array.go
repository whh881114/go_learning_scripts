package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4}

	fmt.Println("1. *** 使用传统的方法：len() ***")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Array at index %d is %d.\n", i, arr[i])
	}

	fmt.Println("2. *** 使用先进的方法：range ***")
	for i, v := range arr {
		fmt.Printf("Array at index %d is %d.\n", i, v)
	}
}
