package main

import (
	"fmt"
)

// 接口(interface)定义了一个对象的行业规范，只定义规范不实现，由具体的对象来实现规范的细节。

// 介绍
// 在Go语言中接口(interface)是一种类型，一种抽象的类型。
// interface是一组method的集合，是duck-type programming的一种休现，接口做的事情就像是定义一个协议（规则），
// 只要一台机器有冼衣服和甩干的功能，我们就称为洗衣机。不关心其属性（数据），只关心行为（方法）。

// 接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer。
// 接口名最好要能突出该接口的类型含义。

type Cat struct{}

func (c Cat) Say() string { return "喵喵喵" }

type Dog struct{}

func (d Dog) Say() string { return "汪汪汪" }

// 使用接口编程
type animal interface {
	Say() string
}

// 只要一个类型它实现了wash()和dry()方法，我们就称这个类型实现了washingMechine这个接口了。
type washingMechine interface {
	wash()
	dry()
}

type Haier struct {
	name  string
	price float32
	mode  string
}

func (h *Haier) wash() {
	fmt.Println("海尔洗衣机能洗衣服。")
}

func (h *Haier) dry() {
	fmt.Println("海尔洗衣机自带甩干功能。")
}

// 田螺姑娘
type tianluo struct {
	name string
}

func (h *tianluo) wash() {
	fmt.Println("田螺姑娘能洗衣服。")
}

func (h *tianluo) dry() {
	fmt.Println("田螺姑娘可以把衣服拧干。")
}

func main() {
	c := new(Cat)
	fmt.Println(c.Say())
	d := new(Dog)
	fmt.Println(d.Say())
	fmt.Println("======================================")

	var animalList []animal
	c1 := new(Cat)
	d1 := new(Dog)
	animalList = append(animalList, c1, d1)
	fmt.Println(animalList)
	for _, i := range animalList {
		fmt.Println(i.Say())
	}
	fmt.Println("======================================")

	var washing_mechine washingMechine
	haier := new(Haier)
	haier.name = "小神童"
	haier.price = 999.99
	haier.mode = "滚洞"

	fmt.Printf("%#v\n", haier)
	fmt.Printf("%#v\n", washing_mechine)

	// Haier结构体实现了washingMechine接口，所以haier这个变量可以赋值给washing_mechine。
	washing_mechine = haier
	fmt.Printf("%#v\n", washing_mechine)

	tl := new(tianluo)
	tl.name = "田螺姑娘"
	washing_mechine = tl
	fmt.Printf("%#v\n", washing_mechine)
}
