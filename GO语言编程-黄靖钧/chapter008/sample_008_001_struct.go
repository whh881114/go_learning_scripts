package main

import (
	"fmt"
	"reflect"
)

type myStruct struct {
	i1  int
	f1  float32
	str string
}

func main() {
	// 初始化的第一种方式，定义后，使用new()给结构体变量分配内存，返回指向已分配内存的指针。
	// ms := new(myStruct)

	// 初始化的第二种方式，注意定义和分配内存两部分代码一定要放在一起。
	// 了解这个方式的好处时，传参给函数时，写的变量形式就是会func(ms *myStruct){}，因为传指针比传值给函数时的代价很小。
	// var ms *myStruct
	// ms = new(myStruct)

	// var ms myStruct，这种声明也是对的，这里完成的操作就是给ms分配了内存，并初始化为零值，
	// 这个时候ms是类型myStruct的一个实例或对象。

	// 变量赋值，适用以上两种方式。
	// ms.i1 = 10
	// ms.f1 = 15.5
	// ms.str = "Google"

	// 初始化的第三种方式，直接赋值
	// 为什么赋值时，后面是&myStruct开头呢，因为ms是初化始化返回的是分配的内存地址。
	// 执行这条语句后，会调用new()初始化，而且大括号里面的值必须按照顺序填写。
	ms := &myStruct{10, 15.5, "Google"}

	// 以上这个方法等同于以下两个语句，这个方式比较好，声明一个myStruct实例，然后给这个实例赋值，简单明了。
	// var ms myStruct
	// ms = myStruct{10, 15.5, "Google"}

	fmt.Printf("int: %d\n", ms.i1)
	fmt.Printf("float: %f\n", ms.f1)
	fmt.Printf("string: %d\n", ms.str)
	fmt.Println("ms data type is: ", reflect.TypeOf(ms), " and the value is: ", ms)
}
