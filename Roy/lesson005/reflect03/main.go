package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s Student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s Student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	// fmt.Println(t.Method())

	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name: %s, method: %s\n", t.Method(i).Name, methodType)
		// var args = []reflect.Value{}
		// v.Method(i).Call(args)
	}
}

func main() {
	stu1 := Student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())
	fmt.Println("========================================================================")

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name: %s, index: %d, type: %v，json_tag: %v\n", field.Name, field.Index, field.Type, field.Tag)
	}
	fmt.Println("========================================================================")

	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name: %s, index: %d, type: %v，json_tag: %v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag)
	}
	fmt.Println("========================================================================")

	var hj = Student{
		Name:  "haojie",
		Score: 100,
	}
	printMethod(hj)
}
