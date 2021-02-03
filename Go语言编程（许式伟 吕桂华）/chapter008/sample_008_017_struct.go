package main

import "fmt"

type HttpResponse struct {
	status_code int
}

func (r *HttpResponse) validResponse() { // r接收的是指向HttpResponse的指针。
	r.status_code = 200
}

func (r HttpResponse) invalidResponse() { // r接收的是指向HttpResponse的实例。
	r.status_code = 500
}

func (r HttpResponse) updateStatus() string { // r接收的是HttpResponse的实例，也就是值类型。在这里在定义其方法。
	return fmt.Sprint(r)
}

func main() {
	var r1 HttpResponse // r1是值，使用这种方式时，r1已被初始化，并且r1结构体中的status_code初始化为0。
	r1.validResponse()
	// r1就是HtppResponse类型，所以他具有validResponse方法。
	// 但是，这个时候要注意了，r1是值类型，执行validResponse函数时，它将传入的是自已的地址，这样才能改变自己
	// status_code变量值，要不然的话就是默认值0，不信的话，那就做个实验。
	fmt.Println(r1.updateStatus())

	r2 := new(HttpResponse) // r2是指针
	r2.validResponse()
	fmt.Println(r2.updateStatus())

	// 实验开始，需要重新声明一个变量。
	var r3 HttpResponse
	r3.invalidResponse() // 因为r3是值类型，所以调用invalidResponse函数时，
	fmt.Printf("实验结果是0就证明自己是对的，直接打印结果，当前实验结果：%d\n", r3.status_code)
	fmt.Println("实验结果是0就证明自己是对的，调用updateStatus函数，当前实验结果：", r3.updateStatus())
}
