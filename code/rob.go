package main

import (
	"code/code/utils"
	"fmt"
)

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	first := nums[0]
	second := utils.Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, utils.Max(first + nums[i], second)
	}
	return second
}

func main() {
	nums := []int{10,7,9,12,6}
	res := rob(nums)
	fmt.Println("res:",res)
}
