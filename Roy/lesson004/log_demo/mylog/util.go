package mylog

import (
	"path"
	"runtime"
)

func getCallerInfo() (fileName, funcName string, line int) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return
	}

	// 根据pc拿到当前执行的函数名
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	fileName = path.Base(file)
	// fmt.Println(funcName, file, line)
	return
}
