package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, t *http.Request) {
	// fmt.Fprint(w, "Hello, world!")
	// w.Write([]byte("Hello, go world!"))

	// 从文件中读取数据写入到w中
	data, err := ioutil.ReadFile("./hello.txt")
	if err != nil {
		fmt.Println("read from file failed, err:", err)
		return
	}
	w.Write(data)
}

func main() {
	// 注册一个处理/函数
	http.HandleFunc("/", sayHello)
	// 启动服务
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		panic(err)
	}
}
