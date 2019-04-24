package main

import (
	"fmt"
	"net/http"
)

const SERVER_PORT = 8080
const SERVER_DOMAIN = "localhost"
const RESPONSE_TEMPLATE = "Hello, world."

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}

func main() {
	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
	http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT),
		"/Users/Kanon/go/src/goTraining/chapter014/ssl.crt",
		"/Users/Kanon/go/src/goTraining/chapter014/ssl.key",
		nil,
	)
}

// 自生成证书
// 1. 生成密钥：openssl genrsa -out ssl.key 2048
// 2. 根据密钥生成证书请求：openssl req -new -key ssl.key -out ssl.csr
// 3. 自签名，生成私有证书：openssl x509 -req -in ssl.csr -signkey ssl.key -out ssl.crt

// 实际运行可以访问，但是页面是404。
