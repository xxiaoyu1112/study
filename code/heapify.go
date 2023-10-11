package main

import "fmt"

// 小根堆 ，倒序排列
func findKthLargest(nums []int, k int) int {
	var heapify func(nums []int, root, end int)
	heapify = func(nums []int, root, end int) {
		for {
			child := root*2 + 1
			if child > end {
				return
			}
			if child < end && nums[child] >= nums[child+1] {
				child++
			}
			if nums[root] < nums[child] {
				return
			}
			nums[child], nums[root] = nums[root], nums[child]
			root = child
		}
	}
	n := len(nums) - 1
	for i := n / 2; i >= 0; i-- {
		heapify(nums, i, n)
	}
	for i := n; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		n--
		heapify(nums, 0, n)
	}
	fmt.Println(nums)
	return nums[len(nums)-k]
}

func main() {
	findKthLargest([]int{4, 2, 5, 3, 6, 7, 2, 1}, 2)
}
