package main

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	// 去重，遍历数组，记录数的存在状态
	for _, num := range nums {
		numSet[num] = true
	}
	// 最长连续序列长度
	longestStreak := 0
	// 判断当前数值的前一个数是否在数组中
	for num := range numSet {
		if !numSet[num-1] {
			// 每次重新记录当前数的值及当前最长连续序列长度
			currentNum := num
			currentStreak := 1
			// 查看后一个数是否在数组中
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}

func main() {
	longestConsecutive([]int{100, 4, 200, 1, 3, 2})
}
