package main

import (
	"container/list"
	"fmt"
)

func main() {
	list := list.New()
	list.PushBack(1)
	list.PushBack(2)

	fmt.Printf("长度：%v\n", list.Len())
	fmt.Printf("第一个元素：%#v\n", list.Front())
	fmt.Printf("第二个元素：%#v\n", list.Front().Next())
}
