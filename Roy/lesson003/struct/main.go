package main

import (
	"fmt"
)

// 结构体，其占用的是一块连续的内存。
// 创建新的类型要使用type关键字
type student struct {
	name   string
	age    int
	gender bool // true，表示男性；false，表示女性。
	hobby  []string
}

func main() {
	// 实例化方法1，student是值类型，不需要初始化，当不赋值时，对应属性值就是其类型的默认值。
	// 我觉得这个方法最好，不需要使用new，指针地址之类的。
	var wanghaohao student
	fmt.Printf("wanghaohao.name is %v\n", wanghaohao.name)
	fmt.Printf("wanghaohao.age is %v\n", wanghaohao.age)
	fmt.Printf("wanghaohao.gender is %v\n", wanghaohao.gender)
	fmt.Printf("wanghaohao.hobby is %v\n", wanghaohao.hobby)
	fmt.Println("======================================================")

	// 实例化方法2
	var roy = new(student)
	fmt.Printf("roy.name is %v\n", roy.name)
	fmt.Printf("roy.age is %v\n", roy.age)
	fmt.Printf("roy.gender is %v\n", roy.gender)
	fmt.Printf("roy.hobby is %v\n", roy.hobby)
	fmt.Println("======================================================")

	// 初始化
	var haojie = student{
		name:   "豪杰",
		age:    18,
		gender: true,
		hobby:  []string{"电影", "看书"},
	}

	fmt.Printf("haojie is %#v\n", haojie)

	// 结构体支持.访问属性。
	fmt.Printf("haojie.name is %v\n", haojie.name)
	fmt.Printf("haojie.age is %v\n", haojie.age)
	fmt.Printf("haojie.gender is %v\n", haojie.gender)
	fmt.Printf("haojie.hobby is %v\n", haojie.hobby)
	fmt.Println("======================================================")
}
