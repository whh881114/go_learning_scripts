package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.1415
	v := reflect.ValueOf(x)
	fmt.Println("通过reflect.ValueOf()获取变量x值后，赋给变量v，检查是否可以被设置：", v.CanSet())

	// 设置一个值：
	// v.SetFloat(3.1415)
	// 上面代码会产生错误，返回：
	// panic: reflect: reflect.Value.SetFloat using unaddressable value

	fmt.Println()
	fmt.Println("*****************************************************************************")
	v = reflect.ValueOf(&x) // 注意：取x的地址，传入是指针就可以设置了。
	fmt.Println("通过reflect.ValueOf()查询x内存地址所在的值类型：", v.Type())
	fmt.Println("将上述得到的结果赋值给v，查询变量v的值是否可以被设置：：", v.CanSet())

	fmt.Println()
	fmt.Println("*****************************************************************************")
	v = v.Elem()
	fmt.Println("v变量是取x地址所对应的值，现在调用的是Elem()函数，返回接口所包含的值：", v)
	fmt.Println("v变量是取x地址所对应的值，现在调用的是Elem()函数，返回接口所包含的值类型：", v.Type())
	fmt.Println("v变量是否可以被设置：", v.CanSet())

	v.SetFloat(3.1415926) // 编译通过
	fmt.Println(v.Interface())
	fmt.Println(v)
}
