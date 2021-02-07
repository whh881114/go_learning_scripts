package main

import "fmt"

// 空接口是指没有定义任何方法的接口，因此任何类型都实现了空接口，所以说定义一个空接口列表，它能存储任意类型。

func showType(x interface{}) {
	// 因为我这个函数可以接收任意类型的变量，类型断言，可以判断变量类型。
	// _, ok := x.(int)
	// if !ok {
	// 	fmt.Println("x不是一个int类型。")
	// } else {
	// 	fmt.Println("x就是一个int类型。")
	// }

	// 升级版
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, value is %v.\n", v)
	case int:
		fmt.Printf("x is a int, value is %v.\n", v)
	case bool:
		fmt.Printf("x is a bool, value is %v.\n", v)
	default:
		fmt.Println("unsupport type!")

	}

}

func main() {
	var x interface{}
	x = 100
	fmt.Println(x)
	showType(x)

	x = "沙河"
	fmt.Println(x)
	showType(x)

	x = false
	fmt.Println(x)
	showType(x)

	x = struct {
		name string
	}{name: "花花"}
	fmt.Println(x)
	showType(x)
	fmt.Println("=======================================")

	stuInfo := make(map[string]interface{}, 100)
	stuInfo["豪杰"] = "名字"
	stuInfo["韩鑫"] = true
	stuInfo["wanghaohao"] = 99.99
	stuInfo["roy"] = [...]string{"a", "b", "c"}
	fmt.Printf("%#v\n", stuInfo)
}
