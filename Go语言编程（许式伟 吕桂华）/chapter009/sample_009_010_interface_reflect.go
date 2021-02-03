package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return fmt.Sprintf("%s - %s - %s", n.s1, n.s2, n.s3)
	// return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = NotknownType{"Ggoogle", "Go", "Code"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret) // 也可以这样写：typ := value.Type()
	knd := value.Kind()

	fmt.Println(typ)
	fmt.Println(knd)

	// 迭代结构的字段
	// NumField()方法返回结构内的字段数量。
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}

	results := value.Method(0).Call(nil)
	fmt.Println(results)
}
