package main

import "fmt"

func shuffle(nums []int, n int) []int {
	ans := make([]int, n*2)
	for i, num := range nums[:n] {
		ans[2*i] = num
		ans[2*i+1] = nums[n+i]
	}
	return ans
}

func main() {
	nums := []int{1,2,3,4,4,3,2,1}
	res := shuffle(nums,4)
	fmt.Println("res:",res)
}
