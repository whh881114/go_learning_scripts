package main

import "fmt"

func main() {
	func(num int) int {
		sum := 0
		for i := 0; i <= num; i++ {
			sum += i
		}
		fmt.Println(sum)
		return sum
	}(100)
}
