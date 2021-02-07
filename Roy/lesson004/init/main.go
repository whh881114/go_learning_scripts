package main

import "fmt"
import _ "github.com/whh881114/go_learning_scripts/Roy/lesson004/package/math_pkg"

// init()示例
var today = "Sunday"

const Week = 7

type student struct {
	Name string
}

// 包被导入的时候会自动执行，多用来做初始化的操作。
func init() {
	fmt.Println("包被初始化了！")
	fmt.Println(today)
}

func main() {
	fmt.Println("这里main函数。")
}
