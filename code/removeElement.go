package main

func removeElement(nums []int, val int) int {
	res := 0
	for _, v := range nums {
		if v != val {
			nums[res] = v
			res++
		}
	}
	nums = nums[:res]
	return res
}

func main() {
	removeElement([]int{3, 2, 2, 3}, 3)
}
