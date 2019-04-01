package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hi, I'm Job, Hi."

	fmt.Printf("The position of \"Job\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Job"))

	fmt.Printf("the position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))

	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Tim\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger")) // -1，表示不存在。

	// 如果是非ASCII编码的字符，可以使用IndexRune函数来对字符进行定位。

	fmt.Println()
	fmt.Println("***************************************************************")
	fmt.Println("如果是非ASCII编码的字符，可以使用IndexRune函数来对字符进行定位。")
	str = "我是左蓝"
	runeStr := []rune(str)
	fmt.Printf("str原始长度为%d。\n", len(str))
	fmt.Printf("str转化为rune后其长度为%d。\n", len(runeStr))
	fmt.Printf("The position of \"左\" is: ")
	fmt.Printf("%d。\n", strings.IndexRune(str, '左'))

	fmt.Println()
	fmt.Println("***************************************************************")
	str = "你好世界，这个世界真好。"
	new := "地球"
	old := "世界"
	n := 1
	fmt.Println(strings.Replace(str, old, new, n)) // 如果n为-1，则表示匹配所有。

	fmt.Println()
	fmt.Println("***************************************************************")
	str = "Golang is cool, right?"
	var manyO = "o"
	fmt.Printf("%d\n", strings.Count(str, manyO))
	fmt.Printf("%d\n", strings.Count(str, "oo"))

	fmt.Println()
	fmt.Println("***************************************************************")
	str = "你好世界"
	fmt.Printf("%d\n", len([]rune(str)))
	fmt.Println(utf8.RuneCountInString(str))

	fmt.Println()
	fmt.Println("***************************************************************")
	orig := "How are you?"
	lower := strings.ToLower(orig)
	upper := strings.ToUpper(orig)

	fmt.Printf("%s\n", orig)
	fmt.Printf("%s\n", lower)
	fmt.Printf("%s\n", upper)

	fmt.Println()
	fmt.Println("***************************************************************")
	fmt.Println("### Trim函数的用法")
	fmt.Printf("%q\n", strings.Trim(" !!! Golang !!! ", "! "))
	fmt.Printf("%q\n", strings.Trim(" !!! Golang !!! ", " ! "))
	fmt.Printf("%q\n", strings.Trim(" !!! Golang !!! ", "!"))

	fmt.Println("### TrimLeft函数的用法")
	fmt.Printf("%q\n", strings.TrimLeft(" !!! Golang !!! ", "! "))
	fmt.Printf("%q\n", strings.TrimLeft(" !!! Golang !!! ", " ! "))
	fmt.Printf("%q\n", strings.TrimLeft(" !!! Golang !!! ", "!"))

	fmt.Println("### TrimSpace函数的用法")
	fmt.Println(strings.TrimSpace(" \t\n这是\t一句话\n\t\r\n"))

	fmt.Println()
	fmt.Println("***************************************************************")
	// Trim函数会将第二个参数转为rune类型，进而逐步替换。
	// 换句话说，实际上"今天"是被分为两个字符分别匹配裁剪的，所以只返回了"气真好"。
	fmt.Printf("%q\n", strings.Trim("今天天气真好", "今"))
	fmt.Printf("%q\n", strings.Trim("今天天气真好", "天"))
	fmt.Printf("%q\n", strings.Trim("今天天气真好", "今天"))

	fmt.Println()
	fmt.Println("***************************************************************")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a boy a girl a dog a cat", "a"))
	fmt.Printf("%q\n", strings.Split("xyz ", ""))

	fmt.Println()
	array1 := strings.Split("a boy a girl a dog a cat", "a")
	fmt.Printf("The type of array1 is %s.\n", reflect.TypeOf(array1))
	for i := 0; i < len(array1); i++ {
		fmt.Printf("Index: %d\t Original Content: %q\n", i, array1[i])
		fmt.Printf("Index: %d\t Trimed Content: %q\n", i, strings.Trim(array1[i], " "))
	}
	// 这个Split函数感觉很变态，第一个a分割后竟然还保留一个空的元素。
}
