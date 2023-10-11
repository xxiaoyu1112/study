package main

import "code/code/utils"

func integerBreak(n int) int {
	dp := make([]int,n+1)
	dp[2] = 1
	for i := 3;i <= n;i++{
		for j := 1;j<i-1;j++{
			dp[i] = utils.Max(dp[i], utils.Max((i - j) * j, dp[i - j] * j))
		}
	}
	return dp[n]
}

func main() {
	integerBreak(10)
}
