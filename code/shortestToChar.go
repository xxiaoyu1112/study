package main

import "fmt"

func ShortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)

	// 从前面开始遍历
	idx := -n
	for i, ch := range s {
		if byte(ch) == c {
			idx = i
		}
		ans[i] = i - idx
	}

	// 从后面开始遍历
	idx = n * 2
	for i := n - 1; i >= 0; i-- {
		a := s[i]
		fmt.Println(a, c)
		if s[i] == c {
			idx = i
		}
		ans[i] = min(ans[i], idx-i)
	}
	return ans
}

//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}

func main() {
	var aa byte = 'e'
	ans := ShortestToChar("loveleetcode", aa)
	fmt.Println("ans:", ans)
}
