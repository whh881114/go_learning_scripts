package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// 只声明没有初始化，a就是初始值nil。
	var a map[string]int
	fmt.Println(a == nil)

	// map初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	// 赋值
	a["沙河娜扎"] = 100
	a["沙河小王子"] = 200
	fmt.Printf("a type is %T, a content is %#v\n", a, a)

	// 声明并初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b type is %T, b content is %#v\n", b, b)

	// 必须要初始化才能赋值
	// var c map[int]int
	// c[100] = 200
	// fmt.Println(c)

	// 判断某个键存不存在
	scoreMap := make(map[string]int, 8)
	scoreMap["沙河娜扎"] = 100
	scoreMap["沙河小王子"] = 200
	scoreMap["沙河小王子2"] = 90
	scoreMap["沙河小王子3"] = 100
	value, ok := scoreMap["张二狗子"]
	if ok {
		fmt.Println("张二狗子在scoreMap中，分数为", value)
	} else {
		fmt.Println("查无此人。")
	}

	// 遍历
	for k, v := range scoreMap {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	for k := range scoreMap {
		fmt.Printf("Key: %s\n", k)
	}

	for _, v := range scoreMap {
		fmt.Printf("Value: %d\n", v)
	}

	// 删除key/value
	delete(scoreMap, "沙河小王子3")
	fmt.Println(scoreMap)
	fmt.Println("=============================================")

	// 按照某个固定顺序遍历map
	var scoreMap2 = make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) // 0-99的随机数
		scoreMap2[key] = value
	}
	for k, v := range scoreMap2 {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}
	fmt.Printf("scoreMap2 content is: %#v", scoreMap2)
	fmt.Println("=============================================")

	// 按照key从小到大的顺序去遍历scoreMap2
	// 先取出所有的key，对key进行排序，最后根据key来取值。
	stuId := []string{}
	for k := range scoreMap2 {
		stuId = append(stuId, k)
	}
	sort.Strings(stuId)
	for i := range stuId {
		fmt.Printf("STUDENT_ID: %s，STUDENT_SCORE: %d。\n", stuId[i], scoreMap2[stuId[i]])
	}
	fmt.Println("=============================================")

	// 复杂的map类型
	var mapSlice = make([]map[string]int, 8, 8) // 数组类型，其元素是map类型
	// mapSlice的值为[nil, nil, nil, nil, nil, nil, nil, nil]
	// 内部的map元素还需要初始化。
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["沙河小王子"] = 100
	fmt.Printf("%#v\n", mapSlice)
	fmt.Println("=============================================")

	var mapSlice2 = make(map[string][]int, 8)
	v, ok := mapSlice2["中国"]
	if ok {
		fmt.Println(v)
	} else {
		mapSlice2["中国"] = make([]int, 8, 8) // 完成了对切片的初始化
		mapSlice2["中国"][0] = 100
		mapSlice2["中国"][1] = 200
		mapSlice2["中国"][2] = 300
	}
	for k, v := range mapSlice2["中国"] {
		fmt.Printf("INDEX: %2d, VALUE: %d\n", k, v)
	}
	fmt.Println("=============================================")

	// 作业：统计一个字符串中每个单词出现的次数。仅限英文字符统计。
	// "the quick brown fox jumps over a lazy dog"
	var exString = "the quick brown fox jumps over a lazy dog"
	var wordMap = make(map[string]int)
	var wordArray = []string{}
	var wordUniq = []string{}

	for i := 0; i < len(exString); i++ {
		word := fmt.Sprintf(string(exString[i]))
		if word == " " {
			continue
		} else {
			wordArray = append(wordArray, word)
		}
	}

	for _, word := range wordArray {
		wordMap[word]++
	}

	for k := range wordMap {
		wordUniq = append(wordUniq, k)
	}
	sort.Strings(wordUniq)

	for _, word := range wordUniq {
		fmt.Printf("字母\"%s\"出现的次数：%d。\n", word, wordMap[word])
	}
}
