package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is an example of a string"

	fmt.Printf("Does the string \"%s\" have prefix %s?\n", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

	fmt.Println()
	fmt.Printf("Does the string \"%s\" have prefix %s?\n", str, "string")
	fmt.Printf("%t\n", strings.HasPrefix(str, "string"))

	fmt.Println()
	fmt.Printf("Does the string \"%s\" have prefix %s?\n", str, "example")
	fmt.Printf("%t\n", strings.Contains(str, "example"))

	fmt.Println()
	fmt.Println("*******************************************************")
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "a & o")) // true
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))

	fmt.Println()
	fmt.Println("*******************************************************")
	fmt.Println(strings.Contains("team", "i"))
	fmt.Println(strings.Contains("failure", "a & o")) // false
	fmt.Println(strings.Contains("foo", ""))          // true
	fmt.Println(strings.Contains("", ""))             // true
}

/* 返回数据：
Does the string "This is an example of a string" have prefix Th?
true

Does the string "This is an example of a string" have prefix string?
false

Does the string "This is an example of a string" have prefix example?
true

*******************************************************
false
true
false
false

*******************************************************
false
false
true
true
*/
