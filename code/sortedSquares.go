package main

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	i, j, k := 0, n-1, n-1
	for i <= j {
		x := nums[i] * nums[i]
		y := nums[j] * nums[j]
		if x < y {
			result[k] = y
			j--
		} else {
			result[k] = x
			i++
		}
		k--
	}
	return result
}

func main() {
	sortedSquares([]int{-7, -3, 2, 3, 11})
}
