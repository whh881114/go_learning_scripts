package main

// 导入包时的路径，要从$GOPATH/src下，写入其相对路径。
import (
	"fmt"
	"github.com/whh881114/go_learning_scripts/Roy/lesson004/package/math_pkg"
)

// 匿名导入包
// 如果只希望导入包，而不使用包内部的数据时，可以使用匿名导入包。
// import _ "包的路径"

func main() {
	sum := math_pkg.Add(1, 2)
	fmt.Println(sum)
}
