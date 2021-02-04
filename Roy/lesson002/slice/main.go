package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3} // 直接声明

	fmt.Printf("a: %T, b: %T\n", a, b)
	fmt.Println(b[1])
	fmt.Printf("b的长度是：%d。\n", len(b))
	fmt.Printf("b的容量是：%d。\n", cap(b))
	fmt.Println("===========================")

	var c []int
	c = a[0:2]
	fmt.Printf("c: %T\n", c)
	fmt.Println("===========================")

	// c = a[:] 从数组的开始切到结束。
	// c = a[:2] 从数组的开始切索引2处，不包含索引2元素。
	// c = a[1:] 从数组的开始切索引1处，包含索引1元素。（左包含，右不包含）

	// 切片的大小（目前元素的数量）
	fmt.Println(len(c))
	fmt.Println("===========================")

	// 切片的容量（底层数组的容量即底层最大能放多少元素）
	x := [...]string{"bj", "sh", "sz", "gz", "cd", "hz", "cq"}
	y := x[1:4]
	fmt.Println(y)
	fmt.Printf("切片y的长度为%d。\n", len(y))
	fmt.Printf("切片y的容量为%d。\n", cap(y))
	fmt.Println("===========================")

	// 切片扩容
	a1 := []int{}
	fmt.Printf("a: %v, len: %d, cap: %d, ptr: %p\n", a1, len(a1), cap(a1), &a)
	for i := 0; i < 10; i++ {
		a1 = append(a1, i)
		fmt.Printf("a: %v, len: %d, cap: %d, ptr: %p\n", a1, len(a1), cap(a1), &a)
	}
	fmt.Println("===========================")

	// copy()，切片是引用类型
	a2 := []int{1, 2, 3}
	b2 := a2
	b2[0] = 100
	fmt.Println(a2, b2)

	a2 = []int{1, 2, 3}
	_ = copy(b2, a2)
	b2[0] = 100
	fmt.Println(a2, b2)
	fmt.Println("===========================")

	// 删除sh，需要使用append来处理。
	a3 := []string{"bj", "sh", "sz", "gz"}
	a3 = append(a3[:1], a3[2:]...)
	fmt.Println(a3)
	fmt.Println("===========================")

	// 定义一个数组
	a4 := [...]int{1, 3, 5, 7, 9, 11, 13}
	b4 := a4[:] // 地址引用传入给切片b，切片底层是基于数组。
	b4[0] = 100
	fmt.Println(a4[0], b4[0])
	fmt.Printf("a4地址：%p\n", &a4) // 去找a4这个变量对应的值的内存地址
	fmt.Printf("b4地址：%p\n", b4)  // 取b4指向的内存地址
	fmt.Println("===========================")

	c4 := a4[2:5]
	fmt.Println(c4, len(c4), cap(c4))

	d4 := c4[:5]
	fmt.Println(d4, len(d4), cap(d4))

	// 扩容政策
	fmt.Printf("a4的容量: %d\n", cap(a4))
	e4 := append(b4, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Printf("a4的容量: %d\n", cap(e4))
	fmt.Println("===========================")

	// make用来给引用类型做初始化，即申请内存空间。
	// new用来创建值类型。
	// var s1 []int //切片声明了，但是不能直接使用。
	// s1[0] = 100 // 不能这样操作
	// s1 = append(s1, 100)，可以使用这样使用。

	s1 := make([]int, 3, 100)
	fmt.Println(s1, len(s1), cap(s1))
}
