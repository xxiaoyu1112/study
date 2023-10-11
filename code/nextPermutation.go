package main

func nextPermutation(nums []int) {
	// 向左遍历，i从倒数第二开始是为了nums[i+1]要存在
	i := len(nums) - 2
	// 寻找第一个小于右邻居的数
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 这个数在数组中存在，从它身后挑一个数，和它换
	if i >= 0 {
		// 从最后一项，向左遍历
		j := len(nums) - 1
		// 寻找第一个大于 nums[i] 的数
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		// 两数交换，实现变大
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 如果 i = -1，说明是递减排列，如 3 2 1，没有下一排列，直接翻转为最小排列：1 2 3
	l, r := i+1, len(nums)-1
	for l < r {
		// i 右边的数进行翻转，使得变大的幅度小一些
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func main() {
	nextPermutation([]int{1, 4, 5, 3})
}
