package main

import "sort"

func intersection(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if res == nil || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return
}

func main() {
	intersection([]int{1, 2, 2, 1}, []int{2, 2})
}
