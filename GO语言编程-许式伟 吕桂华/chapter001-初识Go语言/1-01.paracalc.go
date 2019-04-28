package main

import "fmt"

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}

	resultChan <- sum // 将计算结果发送到channel中
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultChan := make(chan int, 2)

	fmt.Printf("values[:%d] = %v\n", len(values)/2, values[:len(values)/2])
	fmt.Printf("values[%d:] = %v\n", len(values)/2, values[len(values)/2:])

	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)

	// 接收结果, chan是栈，遵守先进后出的原则，所以sum1是计算后5个元素之和的结果。
	sum1, sum2 := <-resultChan, <-resultChan
	fmt.Println("Result:", sum1, sum2, sum1+sum2)
}
