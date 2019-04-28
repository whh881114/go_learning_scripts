package main

import "fmt"

// PersonInfo是一个包含个人详细信息的类型
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

var personDB map[string]PersonInfo

func main() {
	// 需要使用make来创建一个map，如果注释这行，那么这个personDB全是nil，那么下面的赋值语语句就报错了，因为不能向一个nil map
	// 进行赋值操作。

	// 或者换句话说，声明了slice, map和chan后，必须要使用make内建函数进行初始化操作。
	personDB = make(map[string]PersonInfo)

	personDB["12345"] = PersonInfo{"12345", "Tome", "Room 203, ..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101, ..."}

	// 从这个map查找键为"1234"的信息
	person, ok := personDB["1234"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}

	// 元素删除
	delete(personDB, "1")
	person, ok = personDB["1"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1.")
	} else {
		fmt.Println("Did not find person with ID 1.")
	}
}
