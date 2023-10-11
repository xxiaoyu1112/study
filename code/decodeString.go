package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	// 存倍数的栈
	numStack := []int{}
	// 存 待拼接的str 的栈
	strStack := []string{}
	// 倍数的“搬运工”
	num := 0
	// 字符串的“搬运工”
	result := ""
	// 逐字符扫描
	for _, char := range s {
		// 遇到数字
		if char >= '0' && char <= '9' {
			n, _ := strconv.Atoi(string(char))
			// 算出倍数
			num = num*10 + n
			// 遇到 [
		} else if char == '[' {
			// result串入栈
			strStack = append(strStack, result)
			// 入栈后清零
			result = ""
			// 倍数num进入栈等待
			numStack = append(numStack, num)
			// 入栈后清零
			num = 0
			// 遇到 ]，两个栈的栈顶出栈
		} else if char == ']' {
			// 获取拷贝次数
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			// 构建子串
			result = string(str) + strings.Repeat(result, count)
		} else {
			// 遇到字母，追加给result串
			result += string(char)
		}
	}
	return result
}

func main() {
	s := "abc3[cd]xyz"
	res := decodeString(s)
	fmt.Println("res:", res)
}
