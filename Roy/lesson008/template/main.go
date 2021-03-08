package main

import (
	"fmt"
	"io/ioutil"

	// "io/ioutil"
	"html/template"
	// "math/rand"
	"net/http"
	// "strings"
)

type User struct {
	UserName string
	Password string
	Age      int
}

func info(w http.ResponseWriter, r *http.Request) {
	// 第一种访问方法。
	htmlByte, err := ioutil.ReadFile("./info.html")
	if err != nil {
		fmt.Println("Read file failed, err:", err)
		return
	}
	// num := rand.Intn(10)
	// dataStr := string(data)
	// if num > 5 {
	// 	dataStr = strings.Replace(dataStr, "{ooxx}", "<li>《我的世界》</li>", 1)
	// } else {
	// 	dataStr = strings.Replace(dataStr, "{ooxx}", "<li>《小王子》</li>", 1)
	// }
	// w.Write([]byte(dataStr))

	// 添加自定义的方法要在parse模板文件之前添加
	kuaFunc := func(arg string) (string, error) {
		return arg + "真帅", nil
	}

	// 把自定义的函数告诉模板系统
	// template.New("info")                                         // 创建一个template对象
	// template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}) // 给模板系统追加自定义函数
	// 解析模板
	// template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}).Parse(string(htmlByte))

	// 模板文件
	t, err := template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("Read file failed, err:", err)
		return
	}

	// 模板文件
	// t, err := template.ParseFiles("./info.html")
	// if err != nil {
	// 	fmt.Println("Read file failed, err:", err)
	// 	return
	// }

	// data := "<li><<我的世界v2>></li>"
	// t.Execute(w, data)

	// user := User{
	// 	"wanghaohao",
	// 	"123",
	// 	18,
	// }
	// t.Execute(w, user)

	userMap := map[int]User{
		1: {"wanghaohao", "abc", 18},
		2: {"roy", "abc", 16},
		3: {"abc", "abc", 19},
	}
	t.Execute(w, userMap)

}

func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
