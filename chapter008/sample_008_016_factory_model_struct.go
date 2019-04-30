package main

// 不强制使用构造函数，首字母大写
type File struct {
	fd   int    // 文件描述符
	name string // 文件名
}

// 强制使用构造函数，首字母小写
type file struct {
	fd   int    // 文件描述符
	name string // 文件名
}
