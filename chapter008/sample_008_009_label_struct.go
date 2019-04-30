package main

import (
	"fmt"
	"reflect"
)

type TagType struct { // tags，在变量后面加上注释就相当于打标签了
	field1 bool   "是否有存货"
	field2 string "商品名称"
	field3 int    "商品价格"
}

func main() {
	tt := TagType{true, "iPhone X", 8999}

	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

// 在一个变量上调用reflect.TypeOf()可以获取变量的正确类型，
// 如果yojgj一个结构体类型，就可以通过Field来索引结构体的字段，然后就可以使用Tag属性。
func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
