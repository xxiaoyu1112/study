package main

import "fmt"

func missingTwo(nums []int) []int {
	sum := 0
	n := len(nums) + 2
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	a := (1 + n) * n / 2
	sumTwo := a - sum
	limit := sumTwo / 2
	sum = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= limit {
			sum += nums[i]
		}
	}
	b := (1 + limit) * limit / 2
	y := b - sum
	return []int{y, sumTwo - y}
}

func main() {
	res := missingTwo([]int{3, 4})
	fmt.Println("res : ", res)
}
