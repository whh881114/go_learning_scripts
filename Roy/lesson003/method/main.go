package main

import "fmt"

// 函数是谁都可以调用的。

// 方法就是某个具体类型才能调用的方法。

type people struct {
	name   string
	gender string
}

// 函数指定接受者之后就是方法
// func dream() {
// 	fmt.Println("我的梦想是不用上班，也有钱拿！")
// }

// people前面加*，表示引用其原先变量的地址，里面有对struct的属于变更也会影响到后续的操作。
// 不加*也是可以的，只是说方法的目的不一样的而已，两者的写法没有什么错。
// 推荐使用指针类型的接收方式。
func (p *people) dream() {
	p.name = "Roy"
	fmt.Printf("%s的梦想是不用上班，也有钱拿！\n", p.name)
}

func main() {
	// 从这个例子来看，可以根据C++语言做一个类似的理解，go语言中的"类"是有struct定义的属性，func定义了方法。
	p := new(people) // 初始化后，p是一个指针。
	p.name = "wanghaohao"
	p.gender = "male"
	fmt.Printf("Type: %T, Content: %#v\n", p, *p)
	p.dream()
	fmt.Printf("Type: %T, Content: %#v\n", p, *p)
}
