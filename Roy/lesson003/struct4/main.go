package main

import "fmt"

// 模拟继承

type animal struct {
	name string
}

// 定义一个动物会动的方法
func (a *animal) move() {
	fmt.Printf("%s会动。\n", a.name)
}

type dog struct {
	name   string
	feet   int
	animal animal
}

func (d *dog) wangwang() {
	fmt.Printf("%s在叫：汪浩浩。\n", d.name)
}

func main() {
	var a = new(dog)
	a.name = "wanghaohao"
	a.feet = 4
	a.wangwang()

	a.animal.name = "Roy"
	a.animal.move()
}
