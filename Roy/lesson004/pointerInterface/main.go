package main

import "fmt"

// 值接收者和指针接收者实现接口的区别
type animal interface {
	speak()
	move()
}

type animal2 interface {
	speak2()
	move2()
}

type Cat struct {
	name string
	age  int
}

// Cat类型要实现animal的接口
// 值接收者
func (c Cat) speak() {
	fmt.Printf("%s：喵喵喵，它的年龄是%d。\n", c.name, c.age)
	c.age++
}

func (c Cat) move() {
	fmt.Printf("%s：猫猫会动\n", c.name)
}

// 指针接收者
func (c *Cat) speak2() {
	fmt.Printf("%s：喵喵喵，它的年龄是%d。\n", c.name, c.age)
	c.age++
}

func (c *Cat) move2() {
	fmt.Printf("%s：猫猫会动\n", c.name)
}

func main() {
	var x animal
	hh := new(Cat)
	hh.name = "花花"
	hh.age = 2
	x = hh
	x.speak() // 每次都是值copy，所以返回的结果会一直是2。
	x.speak()

	var x2 animal2
	tom := new(Cat)
	tom.name = "Tom猫"
	tom.age = 3
	x2 = tom
	x2.speak2() // 每次都是地址传递，所以返回的结果会变化。
	x2.speak2()

}
