package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	m := map[int]int{0: 0}
	sum := 0
	var allNum []int
	var ans [][]int
	for i := 1; i < target; i++ {
		sum += i
		m[sum] = i
		allNum = append(allNum, i)
		if val, ok := m[sum-target]; ok {
			ans = append(ans, allNum[val:i])
		}
	}
	return ans
}

func main() {
	res := findContinuousSequence(9)
	fmt.Println("res:", res)
}
