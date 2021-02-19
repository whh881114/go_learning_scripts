package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, t *http.Request) {
	data, err := ioutil.ReadFile("./form.html")
	if err != nil {
		fmt.Println("read from file failed, err:", err)
		return
	}
	w.Write(data)
}

func index(w http.ResponseWriter, r *http.Request) {
	// 获取注册信息
	// fmt.Println(r.Method)
	// fmt.Println(r)
	fmt.Printf("%#v\n", r.Form)
	fmt.Printf("%#v\n", r.Form.Get("username"))
	fmt.Printf("%#v\n", r.Form.Get("password"))
	w.Write([]byte("Hello, go world!"))
}

func main() {
	// 注册一个处理/函数
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/index", index)
	// 启动服务
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		panic(err)
	}
}
