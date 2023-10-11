package main

import (
	"fmt"
)

func finalPrices(prices []int) []int {
	n := len(prices)
	ans := make([]int, len(prices))
	st := []int{0}
	for i := n - 1; i >= 0; i-- {
		p := prices[i]
		for len(st) > 1 && st[len(st)-1] > p {
			st = st[:len(st)-1]
		}
		ans[i] = p - st[len(st)-1]
		st = append(st, p)
	}
	return ans
}

func findMin(nums []int) int {
	minIndex := 0
	min := nums[0]
	for i, v := range nums {
		if v < min {
			min = v
			minIndex = i
		}
	}
	return minIndex
}

func main() {
	res := findMin([]int{8, 4, 6, 2, 3})
	fmt.Println("res:", res)
}
