package split

// 定义一个切割字符串的包
// a:b:cd --> [a, b, cd]

// Split 函数是根据给定的分隔符来切割字符串，并且返回一个切片，当前仅支持英文分割。
// func Split(str, sep string) []string {
// 	value := ""
// 	splitString := make([]string, 0)
// 	for i := 0; i < len(str); i++ {
// 		if string(str[i]) != sep {
// 			value += string(str[i])
// 		} else {
// 			splitString = append(splitString, value)
// 			value = ""
// 		}
// 	}
// 	splitString = append(splitString, value)
// 	return splitString
// }

// Split 函数是根据给定的分隔符来切割字符串，并且返回一个切片。
func Split(str, sep string) []string {
	ret := ""
	splitString := make([]string, 0)
	for _, v := range str {
		if string(v) != sep {
			ret += string(v)
		} else {
			splitString = append(splitString, ret)
			ret = ""
		}
	}
	splitString = append(splitString, ret)
	return splitString
}
