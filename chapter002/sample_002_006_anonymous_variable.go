package main

import "fmt"

func GetClass() (stuNum int, className, headTeacher string) {
	return 49, "一班", "张三"
}

func main() {
	stuNum, _, _ := GetClass()
	fmt.Println(stuNum)
}
