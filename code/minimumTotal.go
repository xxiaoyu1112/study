package main

import (
	"code/code/utils"
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	bottom := triangle[len(triangle)-1]
	dp := make([]int, len(bottom))
	for i := range dp {
		dp[i] = bottom[i]
	}

	for i := len(dp) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = utils.Min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func main() {
	t := [][]int{{2},{3,4},{6,5,7},{4,1,8,3}}
	res := minimumTotal(t)
	fmt.Println("res:",res)
}
