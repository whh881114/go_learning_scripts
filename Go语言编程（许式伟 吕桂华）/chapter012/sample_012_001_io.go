package main

import (
	"fmt"
	"sort"
)

// 排序示例

type StuScore struct {
	name  string
	score int
}

type StuScores []StuScore

func (s StuScores) Len() int {
	return len(s)
}

func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"张三", 95},
		{"李四", 91},
		{"赵五", 96},
		{"王认", 90},
	}

	fmt.Println("=========== 默认结果 ===========")
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}
	fmt.Println()

	// StuScores已经实现了sort.Interface接口，看源码内容，可以看出只要实现Len(), Less()和Swap()功能，即完成了该接口。
	// 所以可以使用sort.Sort()和sort.IsSorted()的参数传入。
	sort.Sort(stus)

	fmt.Println("=========== 升序排序结果 ===========")
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}
	fmt.Println("是否已经排序？", sort.IsSorted(stus))

	fmt.Println()
	fmt.Println("=========== 降序排序结果 ===========")
	sort.Sort(sort.Reverse(stus))
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}
	fmt.Println("是否已经排序？", !sort.IsSorted(stus))

	// sort.IsSorted()只支持降序排序结果检查，所以执行Reverse()后，要取值就把原来结果取反。
}
