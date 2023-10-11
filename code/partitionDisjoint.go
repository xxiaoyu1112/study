package main

import "fmt"

func partitionDisjoint(nums []int) int {
	leftmax, curmax := nums[0], nums[0]
	idx := 0
	for i, v := range nums {
		if v < leftmax {
			leftmax = curmax
			idx = i
		} else {
			if v > curmax {
				curmax = v
			}
		}
	}
	return idx + 1
}

func main() {
	nums := []int{3, 0, 2, 4, 1, 5, 6}
	res := partitionDisjoint(nums)
	fmt.Println("res:", res)
}
