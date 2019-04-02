package main

import "log"

func main() {
	test()
}
func test() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("捕获到异常：%v", r)
		}
	}()

	defer func() {
		panic("第二个错误")
	}()

	panic("第一个错误")
}
