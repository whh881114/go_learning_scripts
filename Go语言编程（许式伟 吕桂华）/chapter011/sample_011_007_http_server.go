package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	// 使用net/http包提供的http.ListenAndServer()方法，可以在指定的地址进行监听，开启一个HTTP，在服务器端该方法的原型如下：
	// func ListenAndServer(addr string, handler Handler) error

	// 第一个参数addr即监听地址。
	// 第二个参数表示服务器处理程序，通常为空，这意味着服务器调用http.DefaultServeMux进行处理，
	// 而服务器端编写的业务逻辑处理程序http.Handle()或http.HandleFunc()默认注入http.DefaultServeMuix中，具体代码如下：

	http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	// 自定义http.Server，代码如下：
	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())

}
