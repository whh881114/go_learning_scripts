package main

import "fmt"

var x interface{} // 空接口

func main() {
	x = 1                  // 修改x的值，查看返回结果的变化
	switch i := x.(type) { // 这里表达式只有一句初始化子语句
	case nil:
		fmt.Printf("这里是nil，x的类型是%T。", i)
	case int:
		fmt.Printf("这里是int，x的类型是%T。", i)
	case float64:
		fmt.Printf("这里是float64，x的类型是%T。", i)
	case bool:
		fmt.Printf("这里是bool，x的类型是%T。", i)
	case string:
		fmt.Printf("这里是string，x的类型是%T。", i)
	default:
		fmt.Printf("未知类型。")
	}
}
