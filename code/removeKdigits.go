package main

import "fmt"

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	stack := make([]rune, 0)
	// 遍历num字符串
	for _, c := range num {
		// 只要k>0且当前的c比栈顶的小，则栈顶出栈，k--
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > c {
			stack = stack[:len(stack)-1]
			k--
		}
		// 当前的字符不是"0"，或栈非空（避免0入空栈），入栈
		if c != '0' || len(stack) != 0 {
			stack = append(stack, c)
		}
	}
	// 如果还没删够，要从stack继续删，直到k=0
	for k > 0 && len(stack) != 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	// 如果栈空了，返回"0"，如果栈非空，转成字符串返回
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}

func main() {
	num := "5439219"
	res := removeKdigits(num, 3)
	fmt.Println("res:", res)
}
