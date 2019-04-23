package main

// 处理未知json字符串所需要用的go-simplejson包，以下只是一个示例。

func main() {
	js, err := NewJson([]byte(`{
							"test": {
										"array": [1, "2", 3],
										"int": 10,
										"float": 5.150,
										"bignum": 9999999999999999999,
										"string": "simplejson",
										"bool": true
									}
	`))

	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()
}
