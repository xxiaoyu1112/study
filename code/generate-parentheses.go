package main

import "fmt"

func generateParenthesis(n int) []string {
	res := []string{}

	var dfs func(lRemain int, rRemain int, path string)
	// 左右括号所剩的数量，str是当前构建的字符串
	dfs = func(lRemain int, rRemain int, path string) {
		// 字符串构建完成
		if 2*n == len(path) {
			// 加入解集
			res = append(res, path)
			// 结束当前递归分支
			return
		}
		// 只要左括号有剩，就可以选它，然后继续做选择（递归）
		if lRemain > 0 {
			dfs(lRemain-1, rRemain, path+"(")
		}
		// 右括号比左括号剩的多，才能选右括号
		if lRemain < rRemain {
			// 然后继续做选择（递归）
			dfs(lRemain, rRemain-1, path+")")
		}
	}
	// 递归的入口，剩余数量都是n，初始字符串是空串
	dfs(n, n, "")
	return res
}

func main() {
	n := 3
	res := generateParenthesis(n)
	fmt.Println("res:",res)
}
