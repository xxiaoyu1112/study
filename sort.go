package main

import (
	"fmt"
)

func sortArray(nums []int) []int {
	var quickSort func(nums []int, l, r int) []int
	quickSort = func(nums []int, l, r int) []int {
		if l >= r {
			return nums
		}
		i, j, cur := l, r, nums[l]
		for i < j {
			for i < j && nums[j] >= cur {
				j--
			}
			for i < j && nums[i] <= cur {
				i++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[i], nums[l] = nums[l], nums[i]
		quickSort(nums, l, i-1)
		quickSort(nums, i+1, r)
		return nums
	}
	return quickSort(nums, 0, len(nums)-1)
}

func heapSortArray(nums []int) []int {
	var heapify func(nums []int, root, end int)
	heapify = func(nums []int, root, end int) {
		for {
			child := root*2 + 1
			if child > end {
				break
			}
			if child < end && nums[child] < nums[child+1] {
				child++
			}
			if nums[root] > nums[child] {
				return
			}
			nums[root], nums[child] = nums[child], nums[root]
			root = child
		}
	}
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		heapify(nums, i, end)
	}
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		heapify(nums, 0, end)
	}
	return nums
}

func main() {
	res := heapSortArray([]int{2, 1, 35, 61, 22, 23, 12})
	fmt.Println(res)
}
