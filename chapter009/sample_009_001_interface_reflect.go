package main

import "fmt"

type Shaper interface {
	// 可不可以这样理解：接口是把"对象"所有的方法都声明在他在接口里了，当调用时，直接把接口赋给一个变量，然后
	// 接口这个变量直接来调对应的函数就可以了？
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	/*
		定义areaIntf
		areaIntf = sq1

		更简洁，不需要分开定义：
		areaIntf := Shaper(sq1)

		甚至这样：
	*/

	fmt.Println("这个执行方法实质是，调用了sq1这个Square实例的方法而已。")
	areaInft := sq1
	fmt.Printf("面积为：%f\n", areaInft.Area())

	fmt.Println("**************************************************")
	fmt.Println("现在使用接口来计算结果。")
	ret := Shaper.Area(sq1)
	fmt.Println("面积为：", ret)

	// 现在知道怎么用了吧，这个接口类对fmt包中的Println接口思想即可，你打印的内容就是相当于这里定义的sq1类似。
	// str1 = "Hello, world!"
	// fmt.Println(str1)
	//
	// 	ret := Shaper.Area(sq1)
	// 	fmt.Println("面积为：", ret)

}
