package main

import "fmt"

// 自己实现一个构造函数

type student struct {
	name   string
	age    int
	gender bool // true，表示男性；false，表示女性。
	hobby  []string
}

// 此时结构体构造函数返回的是一个结构体值类型，也就是说，每调用一次，都需要把值拷贝一份，当结构体中的字段达到几十个时，并且调用频繁，
// 那么内存消耗比较大，所以返回一个指针即可，来节约内存。
// func newStudent(name string, age int, gender bool, hobby []string) student {
// 	return student{
// 		name:   name,
// 		age:    age,
// 		gender: gender,
// 		hobby:  hobby,
// 	}
// }

func newStudent(name string, age int, gender bool, hobby []string) *student {
	return &student{
		name:   name,
		age:    age,
		gender: gender,
		hobby:  hobby,
	}
}

func main() {
	haojie := newStudent("豪杰", 18, true, []string{"电影"})
	fmt.Printf("%v\n", *haojie)
	fmt.Printf("%#v\n", *haojie)
}
