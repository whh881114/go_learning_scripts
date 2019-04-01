package main

import "fmt"

func main() {
	sum := 11
	count := 3
	mean := float32(sum) / float32(count)
	mean2 := sum / count
	fmt.Printf("mean的值为：%f\n", mean)
	fmt.Printf("mean2的值为：%d\n", mean2)
}
