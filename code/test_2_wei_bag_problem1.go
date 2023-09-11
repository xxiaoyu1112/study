package main

import (
	"code/code/utils"
	"fmt"
)

func test2WeiBagProblem1(weight, value []int, bagWeight int) int {
	// 定义dp数组、dp[i][j] 表示从下标为[0-i]的物品里任意取，放进容量为j的背包，价值总和最大是多少。
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}
	// 初始化
	for j := bagWeight; j >= weight[0]; j-- {
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}
	// 递推公式
	for i := 1; i < len(weight); i++ {
		//正序,也可以倒序
		for j := 0; j <= bagWeight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = utils.Max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagWeight]
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	res := test2WeiBagProblem1(weight, value, 4)
	fmt.Println("res:", res)
}
