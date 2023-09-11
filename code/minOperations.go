package main

import (
	"code/code/utils"
	"fmt"
)

func minOperations(nums []int, x int) int {
	target := -x
	for _, x := range nums {
		target += x
	}
	if target < 0 { // 全部移除也无法满足要求
		return -1
	}
	ans, left, sum := -1, 0, 0
	for right, x := range nums {
		sum += x
		for sum > target { // 缩小子数组长度
			sum -= nums[left]
			left++
		}
		if sum == target {
			ans = utils.Max(ans, right-left+1)
		}
	}
	if ans < 0 {
		return -1
	}
	return len(nums) - ans
}

func main() {
	num := []int{1, 1, 4, 2, 3}
	res := minOperations(num, 5)
	//fmt.Println(reflect.TypeOf(len(num)))
	//fmt.Println(reflect.TypeOf(len(num[:])))
	fmt.Println(res)
}
