package main

import (
	"code/code/utils"
	"fmt"
)

// testCompletePack1 先遍历物品, 在遍历背包
func testCompletePack1(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	for i := 0; i < len(weight); i++ {
		// 正序会多次添加 value[i]
		for j := weight[i]; j <= bagWeight; j++ {
			// 推导公式
			dp[j] = utils.Max(dp[j], dp[j-weight[i]]+value[i])
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}

// testCompletePack2 先遍历背包, 在遍历物品
func testCompletePack2(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	// j从0 开始
	for j := 0; j <= bagWeight; j++ {
		for i := 0; i < len(weight); i++ {
			if j >= weight[i] {
				// 推导公式
				dp[j] = utils.Max(dp[j], dp[j-weight[i]]+value[i])
			}
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}

func main() {
	weight := []int{1, 3, 4}
	price := []int{15, 20, 30}
	fmt.Println(testCompletePack1(weight, price, 4))
	fmt.Println(testCompletePack2(weight, price, 4))
}
