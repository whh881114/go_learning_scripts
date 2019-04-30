package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "https://studygolang.com/"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: 请求网站\"%s\"出现错误，现在退出。", url)
		os.Exit(1)
	}

	fmt.Printf("%T\n", resp)
	fmt.Printf("%s\n", resp)

	/*
		*http.Response
		&{200 OK %!s(int=200) HTTP/1.1 %!s(int=1) %!s(int=1) map[Connection:[keep-alive] Content-Type:[text/html; charset=utf-8] Date:[Sun, 21 Apr 2019 01:38:13 GMT] Server:[nginx] Set-Cookie:[user=MTU1NTgxMDY5M3xEdi1CQkFFQ180SUFBUkFCRUFBQUp2LUNBQUVHYzNSeWFXNW5EQXNBQ1VsT1JFVllYMVJCUWdaemRISnBibWNNQlFBRFlXeHN8KvMmcuj5B79tV5HQDgfWw2iuPV0vI1opuQXB6EunibE=; Path=/; Expires=Tue, 21 May 2019 01:38:13 GMT; Max-Age=2592000; HttpOnly] Vary:[Accept-Encoding] X-Request-Id:[cb4663ca-806c-46b0-9c68-bb193684bf36]] %!s(*http.gzipReader=&{0xc000034080 <nil> <nil>}) %!s(int64=-1) [chunked] %!s(bool=false) %!s(bool=true) map[] %!s(*http.Request=&{GET 0xc0000cc000 HTTP/1.1 1 1 map[] <nil> <nil> 0 [] false studygolang.com map[] map[] <nil> map[]   <nil> <nil> <nil> <nil>}) %!s(*tls.ConnectionState=&{771 true false 47 http/1.1 true  [0xc0000ba580 0xc0000bab00] [[0xc0000bb080 0xc0000bb600 0xc0000bbb80]] [] [] 0x5a7c00 [215 74 124 24 120 117 34 176 233 246 60 126]})}
	*/
}
