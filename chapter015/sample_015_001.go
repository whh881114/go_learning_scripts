package main

func toHeap() *int {
	var x int
	return &x
}

func toStack() int {
	x := new(int)
	*x = 1
	return *x
}

func main() {

}

// 执行命令：go run -gcflags '-m -l' sample_015_001.go
/*
	由于栈的性能要比堆好很多，分配内存的速度比堆快（小内存分配），所以在编码过程中，开发人员会比较在意一个变量或者对象被分配到哪里，
	是堆还是栈这对于程序的性能和安全都有较大影响。

	为了显示上面两个函数的区别，需要开启逃逸分析日日土已，在编译的时候加上-gcflags '-m'即可，不过为了不让编译器自动内连函数，一般
	会加-l参数，所以最终编译参数为-gcflags '-m -l'，-m意思是打印编译器的自动优化策略。
*/
