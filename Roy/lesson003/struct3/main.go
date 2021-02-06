package main

import "fmt"

// 匿名字段，把类型当成变量名
type student struct {
	name   string
	string // 类型不同存在相同的
	int
}

// 嵌套结构体，也可以嵌套匿名结构体。
type address struct {
	province string
	city     string
}

type student2 struct {
	name string
	age  int
	addr address //嵌套结构体
}

func main() {
	stu1 := new(student)
	stu1.name = "wanghaohao"
	stu1.string = "male"
	stu1.int = 32
	fmt.Println(stu1)

	stu2 := new(student2)
	stu2.name = "Roy"
	stu2.age = 32
	stu2.addr.province = "GZ"
	stu2.addr.city = "SZ"
	fmt.Println(stu2)

}
