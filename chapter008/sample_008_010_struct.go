package main

import "fmt"

type firstS struct {
	in1 int
	in2 int
}

// 匿名字段和面向对象编程中的继承概念相似，可以被用来模拟类似继承的行为。
// Go语言中的继随是通过内嵌或组合来实现的，所以可以说在Go语言中，组合比继承更受欢迎。
type secondS struct {
	b      int
	c      float32
	int    // 匿名字段
	firstS // 匿名字段
}

func main() {
	sec := new(secondS)
	sec.b = 6
	sec.c = 7.5
	sec.int = 60
	sec.in1 = 5
	sec.in2 = 10

	fmt.Printf("sec.b is : %d\n", sec.b)
	fmt.Printf("sec.c is : %f\n", sec.c)
	fmt.Printf("sec.int is : %d\n", sec.int)
	fmt.Printf("sec.in1 is : %d\n", sec.in1)
	fmt.Printf("sec.in2 is : %d\n", sec.in2)

	// 使用结构体字面量
	sec2 := secondS{6, 7.5, 60, firstS{5, 10}}
	fmt.Println("sec2 is:", sec2)
}
