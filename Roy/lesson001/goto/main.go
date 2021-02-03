package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 2 {
				goto breakTag
			}
			fmt.Printf("%d -- %d\n", i, j)
		}
	}

breakTag:
	fmt.Println("goto跳出两层for循环。")
}
