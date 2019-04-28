package main

import "fmt"

func main() {
LOOP1:
	for {
		x := 1
		switch {
		case x > 0:
			fmt.Println("A")
			break LOOP1 // 跳出标签为LOOP1的代码块外。
		case x == 1:
			fmt.Println("B")
		default:
			fmt.Println("C")
		}
	}
}
