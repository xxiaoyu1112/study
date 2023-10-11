package main

import "fmt"

//func maxSubArray(nums []int) int {
//	max := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i]+nums[i-1] > nums[i] {
//			nums[i] += nums[i-1]
//		}
//		if nums[i] > max {
//			max = nums[i]
//		}
//	}
//	return max
//}

func maxSubArray(nums []int) []int {
	n := len(nums)
	// 这里的dp[i] 表示，最大的连续子数组和，包含num[i] 元素
	dp := make([]int, n)
	// 初始化，由于dp 状态转移方程依赖dp[0]
	dp[0] = nums[0]
	// 初始化最大的和
	mx := nums[0]
	l, r := 0, 0
	for i := 1; i < n; i++ {
		// 这里的状态转移方程就是：求最大和
		// 会面临2种情况，一个是带前面的和，一个是不带前面的和
		if dp[i-1]+nums[i] >= nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
			l = i
		}
		if mx < dp[i] {
			mx = dp[i]
			r = i
		}
	}
	fmt.Println(mx)
	fmt.Println(l, r)
	if l == len(nums)-1 || l > r {
		if nums[r] <= 0 {
			fmt.Println(nums[r])
			return []int{nums[r]}
		} else {
			fmt.Println(nums[0 : r+1])
			return nums[:r+1]
		}
	} else {
		fmt.Println(nums[l : r+1])
		return nums[l : r+1]
	}
	//fmt.Println(mx)
	//return mx
}

func main() {
	maxSubArray([]int{3, 2, -3, -1, 1, -3, 1, -1})
	//fmt.Println("maxSubArray:", max)
}
