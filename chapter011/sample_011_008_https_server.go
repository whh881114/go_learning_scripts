package main

import "net/http"

func main() {
	// net/http包还提供http.ListenAndServeTLS()方法，用于处理HTTPS连接请求：
	func http.ListenAndServeTLS(addr string, certFile string, keyFile string, handler Handlere) error

	// https的请求方式与http一样，除了服务器上必须配置证书和对应的私钥文件外。
}
