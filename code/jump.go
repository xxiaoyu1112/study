package main

import (
	"code/code/utils"
	"fmt"
)

func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length - 1; i++ {
		maxPosition = utils.Max(maxPosition, i + nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}


func main() {
	nums := []int{2,3,1,1,4}
	step := jump(nums)
	fmt.Println("step:",step)
}
