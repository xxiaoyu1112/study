package main

import (
	"code/code/utils"
	"fmt"
)

// 贪心
//func wiggleMaxLength(nums []int) int {
//	n := len(nums)
//	if n < 2 {
//		return n
//	}
//	ans := 1
//	prevDiff := nums[1] - nums[0]
//	if prevDiff != 0 {
//		ans = 2
//	}
//	for i := 2; i < n; i++ {
//		diff := nums[i] - nums[i-1]
//		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
//			ans++
//			prevDiff = diff
//		}
//	}
//	return ans
//}

// 动态规划
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up, down := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = utils.Max(up, down+1)
		} else if nums[i] < nums[i-1] {
			down = utils.Max(up+1, down)
		}
	}
	return utils.Max(up, down)
}

func main() {
	nums := []int{1,2,3,4}
	res := wiggleMaxLength(nums)
	fmt.Println("res:",res)
}
