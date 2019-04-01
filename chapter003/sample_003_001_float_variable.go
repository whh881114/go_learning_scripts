package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value1 float64
	value1 = 1
	value2 := 2
	value3 := 3.0

	// v := value1 + value2 // 编译失败，value1与value2类型不同

	v := value1 + value3

	fmt.Println(value1, value2, value3, v)
	fmt.Println("v的类型是：", reflect.TypeOf(v))
}

// 使用了反射（reflect）的概念返回变量的类型，修改TypeOf中的变量名称。
