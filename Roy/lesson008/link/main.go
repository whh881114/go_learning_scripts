package main

import "fmt"

type student struct {
	name string
}

func (s student) learn() student {
	fmt.Printf("%s热爱学习！\n", s.name)
	return s
}

func (s student) doHomework() {
	fmt.Printf("%s热爱交作业！\n", s.name)
}

func main() {
	haojie := student{
		"豪杰",
	}

	haojie.learn().doHomework()

}
