package main

import "fmt"

func main() {
	grade := "B"
	marks := 90

	switch marks { // switch后面接变量名，case就是实际值。
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Printf("你的成绩为%s\n", grade)

	fmt.Println("*****************************************")
	switch { // switch后面不接变量名，在case中进行逻辑运算。
	case marks >= 90:
		grade = "A"
	case marks >= 80:
		grade = "B"
	case marks >= 70:
		grade = "C"
	case marks >= 60:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Printf("你的成绩为%s\n", grade)

	fmt.Println("*****************************************")
	switch {
	case grade == "A":
		fmt.Println("成绩优秀！")
	case grade == "B":
		fmt.Println("表现良好！")
	case grade == "C", grade == "D": // case表达式可以有多个
		fmt.Println("再接再厉！")
	// 以上的写法，等同于fallthrough
	// case grade == "C":
	//		fallthrough		// 将当前case控制权交张下一个case语句判断。
	// case grade == "D":
	// 	fmt.Println("再接再厉！")
	default:
		fmt.Println("成绩不合格！")
	}
	fmt.Printf("你的成绩为%s\n", grade)
}
