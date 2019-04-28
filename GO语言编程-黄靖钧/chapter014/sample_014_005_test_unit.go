package main

import "fmt"

func main() {
	fmt.Println(Age(-7))
}

// Age测试
func Age(n int) int {
	if n > 0 {
		return n
	}
	n = 0
	return n
}
