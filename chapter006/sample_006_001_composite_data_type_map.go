package main

import "fmt"

func main() {
	// 创建一个映射，存储颜色以及颜色对应的十六进制代友
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7f50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s value: %s\n", key, value)
	}

	fmt.Println("****************************************")
	// 调用函数来移除指定的键
	removeColor(colors, "Coral")

	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s value: %s\n", key, value)
	}
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
