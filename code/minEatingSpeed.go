package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, maxArray(piles)
	helper := func(x int) bool {
		s := 0
		for _, p := range piles {
			s += (p + x - 1) / x
		}
		return s <= h
	}
	for left < right {
		mid := (left + right) >> 1
		if helper(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func maxArray(nums []int) int {
	ans := nums[0]
	for _, num := range nums {
		if num > ans {
			ans = num
		}
	}
	return ans
}

func main() {
	pile := []int{3,6,7,11}
	res := minEatingSpeed(pile,8)
	fmt.Println("res:",res)
}
