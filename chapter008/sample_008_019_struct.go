package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}

func main() {
	p := new(Person)
	// p.firstName 未定义返回错误，因为new只是分配了内存地址，并没有对其实始化。

	p.SetFirstName("Tom")
	fmt.Println(p.FirstName())
}
