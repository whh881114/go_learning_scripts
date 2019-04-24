package main

import "net/http"

func main() {
	h := http.FileServer(http.Dir("."))
	http.ListenAndServeTLS(":8001", "rui.crt", "rui.key", h)
}

// 伪代码，单张介绍了FileServer函数的存在，这个函数的参数只有一个就是根目录路径，其他任务全部由ListenAndServe*等函数完成。
