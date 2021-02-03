package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符器操作
	s1 := "alexdsb"
	fmt.Println(len(s1))

	// 字符串拼接
	s2 := "Python"
	fmt.Println(s1 + s2)

	s3 := fmt.Sprintf("%s---------%s", s1, s2)
	fmt.Println(s3)

	s4 := s3 + "-----wanghaohao"
	fmt.Println(s4)

	ret := strings.Split(s4, "-")
	fmt.Println(ret)

	fmt.Println(strings.Contains(s1, "sb"))
	fmt.Println(strings.HasPrefix(s1, "alx"))
	fmt.Println(strings.HasSuffix(s1, "sb"))
	fmt.Println(strings.Index(s4, "h"))

	myArray := []string{"shell", "perl", "python"}
	fmt.Println(strings.Join(myArray, "-"))
}
