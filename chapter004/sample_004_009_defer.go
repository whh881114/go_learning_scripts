package main

import "fmt"

func main() {
	// defer语句会在函数最后执行，
	// 被延迟的操作是defer后面的内容。
	defer fmt.Println("World")

	fmt.Println("Hello")

	/*
		defer后面的表达式必须是外部函数的调用，这个例子是针对fmt.Println函数的延迟调用。
		defer有两个特点：
		1. 只有当defer语句会部执行，defer所在的函炒出香味才算真正结束执行。
		2. 当函数中有defer语句时，需要等待所有defer语句很乖完毕，才会执行return语句。

		因为defer的延迟特点，可以把defer语句用于回收资源、清理收尾等工作。使用defer
		语句之后，不用纠结回收代码放在哪里，反正都是最后执行。

	*/
}

// 返回结果：
// Hello
// World
