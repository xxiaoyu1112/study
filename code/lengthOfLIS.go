package main

import (
	"code/code/utils"
	"fmt"
)

func lengthOfLIS(nums []int) []int {
	dp := []int{}
	for _, num := range nums {
		if len(dp) == 0 || dp[len(dp)-1] < num {
			dp = append(dp, num)
		} else {
			l, r := 0, len(dp)-1
			pos := r
			for l <= r {
				mid := (l + r) / 2
				if dp[mid] >= num {
					pos = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			dp[pos] = num
		} //二分查找
	}
	return dp
}

func lengthOfLIS2(nums []int) []int {
	var stack []int
	for _, num := range nums {
		l := 0
		// 二分查找当前元素 num 比栈中元素小的位置， 如果存在 num 比栈中某元素小，直接替换栈中的当前元素
		for r := len(stack); l < r; {
			mid := (l + r) / 2
			if stack[mid] >= num {
				r = mid
			} else {
				l = mid + 1
			}
		}
		// 如果 num 比栈中数大的话，l 位置等于数组长度，此时num直接进栈，否则替换l位置的元素为num
		if l == len(stack) {
			stack = append(stack, num)
		} else {
			stack[l] = num
		}
	}
	return stack
}

// 动态规划求解
func lengthOfLIS3(nums []int) int {
	// dp数组的定义 dp[i]表示取第i个元素的时候，表示子序列的长度，其中包括 nums[i] 这个元素
	dp := make([]int, len(nums))

	// 初始化，所有的元素都应该初始化为1
	for i := range dp {
		dp[i] = 1
	}

	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = utils.Max(dp[i], dp[j]+1)
			}
		}
		fmt.Println("dp", dp)
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func main() {
	//nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	nums := []int{0, 1, 0, 3, 2}
	//dp1 := lengthOfLIS(nums)
	stack := lengthOfLIS2(nums)
	//fmt.Println("dp1:",dp1)
	fmt.Println("stack:", stack)
}
