package main

import (
	"encoding/json"
	"fmt"
)

// json序列化

type Student struct { // 首字母大写
	ID     int    `json:"id"` // json的:后面不能带有空格。
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

func main() {
	stu1 := new(Student)
	stu1.ID = 1
	stu1.Name = "wanghaohao"
	stu1.Gender = "male"

	// 序列化：把编程语方里面的数据转换成json格式的字符串
	v, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("JSON格式化出错了。")
		fmt.Println(err)
	}

	fmt.Println(v)                 // []byte
	fmt.Println(string(v))         // 把[]byte转换成string
	fmt.Printf("%#v\n", string(v)) // 把[]byte转换成string

	// 反序列化：把满足json格式的字符串转换成当前编程语言里的对象
	// str := "{\"ID\":1,\"Gender\":\"male\",\"Name\":\"Roy\"}"
	str := "{\"id\":1,\"gender\":\"male\",\"name\":\"Roy\"}" // 其key可以为原来的struct中定义的或是tag指定的名字。
	stu2 := new(Student)
	json.Unmarshal([]byte(str), stu2)
	fmt.Printf("Type: %T, Content: %#v\n", stu2, *stu2)
}
