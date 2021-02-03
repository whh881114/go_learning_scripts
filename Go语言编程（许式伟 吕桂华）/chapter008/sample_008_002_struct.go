package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) { // 接收的是指针变量
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 作为值类型，就是类似var intSample int一样，只是这里的变量类型是struct而已。
	var pers1 Person
	pers1.firstName = "张"
	pers1.lastName = "三三"
	upPerson(&pers1) // 需要传的是指针，因为person是一个值类型，所以前面加个&来取这个值的地址。
	fmt.Printf("这个人的名字是%s %s。\n", pers1.firstName, pers1.lastName)

	// 作为指针，使用new()初始化后，返回的是内存地址，所以在使用函数调用时就可以直接使用这个变量，不需要在前面加上&。
	pers2 := new(Person)
	pers2.firstName = "Roy"
	pers2.lastName = "Wong"
	upPerson(pers2)
	fmt.Printf("这个人的名字是%s %s。\n", pers2.firstName, pers2.lastName)

	// 作为字面量，使用这个方法初始化时，会调用new()来初始化，所以pers3也是一个内存地址。
	pers3 := &Person{"Roy", "Wong"}
	upPerson(pers3)
	fmt.Printf("这个人的名字是%s %s。\n", pers3.firstName, pers3.lastName)

	// 使用这个方式来声明并初始化，这个是作为值类型，所以在调函数时，需要在前面加上&来取地址。
	pers4 := Person{"Roy", "Wong"}
	upPerson(&pers4)
	fmt.Printf("这个人的名字是%s %s。\n", pers4.firstName, pers4.lastName)
}
