package main

import "fmt"

type S struct {
	Name string
}

func (s S) M1() {
	s.Name = "value"
}

func (s *S) M2() {
	s.Name = "pointer"
}

func main() {
	var s1 = S{"new"} // 值类型
	var s2 = &s1      // s1地址赋给s2
	s1.M2()
	fmt.Printf("%+v, %+v\n", s1, s2) // 当s1改变Name原先的值为pointer，s2也做了改变。

	s1 = S{"new"}
	s2 = &s1
	s2.M1() // 当s2改变Name原先的值为pointer，s1也做了改变。
	fmt.Printf("%+v, %+v\n", s1, s2)
}
