package main

import "fmt"

func subarraySum(nums []int, k int) int {
	count := 0
	hash := map[int]int{0: 1}
	preSum := 0

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if hash[preSum-k] > 0 {
			count += hash[preSum-k]
		}
		hash[preSum]++
	}
	return count
}

func main() {
	nums := []int{1,2,3}
	k := 2
	res := subarraySum(nums,k)
	fmt.Println("res:",res)
}
