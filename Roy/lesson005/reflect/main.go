package main

// 反射：变量的类型太多了，类型断言猜不全，使用反射就能直接拿到接口值的动态类型和动态值。
// 反射的应用：各种web框架，配置文件解析库，ORM框架。
// 反射是一把双刃剑，优点是让代码更灵活，缺点是运行效率低。

// 在Go语言的反射机制中，任何接口值都由一个具体类型和具体类型值两部分组成的。在Go语言中反射的相关功能是由内置的reflect包提供，
// 任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，并且reflect.TypeOf和reflect.ValueOf两个函数来
// 获取任意对象的Type和Value。

import (
	"fmt"
	"reflect"
)

type cat struct {
	name string
}

type person struct {
	name string
	age  uint8
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("Type: %v\n", v)
	fmt.Printf("Type: %T\n", x) // 本质是反射。
	fmt.Printf("name: %v, kind: %v\n", v.Name(), v.Kind())
}

func main() {
	reflectType(100)
	reflectType(false)
	reflectType("沙河有沙又有河")
	reflectType([3]int{1, 2, 3})
	reflectType(map[string]int{})

	c := new(cat)
	whh := new(person)
	reflectType(c)
	reflectType(whh)
}
