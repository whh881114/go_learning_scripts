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
	feet int
	// animal animal <==> animal ，如果使用这两种写法，那么只会继承到animal的方法，访问不到animal的属性。所以这个就是值引用。
	// 后面他要调用其animal函数时，不需要对animal属性进行初始化。

	// 像这样写，就是引用了animal类型，dog这个struct就是继承animal的属性和方法。此后，要使用animal变量的时候，需要对它进行初始化。
	// a.animal = new(animal)，肯定要使用这个写法了。
	*animal
}

func (d *dog) wangwang() {
	fmt.Printf("%s在叫：汪浩浩。\n", d.name)
}

func main() {
	var a = new(dog)
	a.animal = new(animal)
	a.name = "wanghaohao"
	a.feet = 4
	a.wangwang()
	a.animal.move()
}
