package main

import "fmt"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	hash := make([]int, 128)
	// 记录子串的各个字符出现次数
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}
	l, count, max, results := 0, len(t), len(s)+1, ""
	for r := 0; r < len(s); r++ {
		// 将遍历到的元素减 1 记录到哈希表 hash 中。
		hash[s[r]]--
		// 此时只有 t 中出现过的字符的hash值才会大于等于 0,说明是 t 中元素，count 减一。
		if hash[s[r]] >= 0 {
			count--
		}
		// 此时 s[l] 这个字符在主串中出现过 但不是子串内的字符或者子串中已经有该字符， 因此 l++ ，判断下一个元素是否为子串所需元素
		for l < r && hash[s[l]] < 0 {
			// 出现次数 加 1
			hash[s[l]]++
			// 左指针向后移动
			l++
		}
		// 此时有符合要求子字符串, 开始缩小滑动窗口。
		if count == 0 && max > r-l+1 {
			max = r - l + 1
			results = s[l : r+1]
		}
	}
	return results
}

func main() {
	s := "ADOBECODEBAANC"
	t := "ABC"
	res := minWindow(s, t)
	fmt.Println(res)
}
