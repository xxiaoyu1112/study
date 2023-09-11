package main

func uniqueLetterString(s string) int {
	n := len(s)
	memo := [26]int{}
	// 初始时，预处理假设所有字符都在 -1 出现过，则后续计算贡献时无需其他判断
	for i := range memo {
		memo[i] = -1
	}

	// 找左侧最近相同字符下标，计算子字符串的可选左端点
	left := make([]int, n)
	for i := 0; i < n; i++ {
		c := s[i]-'A'
		left[i] = i-memo[c]
		memo[c] = i
	}

	// 初始时，预处理假设所有字符都在 n 出现过，则后续计算贡献时无需其他判断
	for i := range memo {
		memo[i] = n
	}

	// 找右侧最近相同字符下标，计算子字符串的可选右端点
	right := make([]int, n)
	for i := n-1; i >= 0; i-- {
		c := s[i]-'A'
		right[i] = memo[c]-i
		memo[c] = i
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += left[i] * right[i]
	}
	return ans
}

func main() {
	uniqueLetterString("ABC")
}
