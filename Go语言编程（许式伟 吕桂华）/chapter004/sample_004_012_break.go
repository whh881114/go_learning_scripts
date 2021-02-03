package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i的值是：%d\n", i)
		if i > 4 {
			break
		}
	}
}
