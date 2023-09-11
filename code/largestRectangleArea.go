package main

import "code/code/utils"

func largestRectangleArea(heights []int) int {
	// 目标找每个柱子左右两边第一个小于该柱子的柱子
	// 单调递减栈
	st := []int{0}
	// 数组头部加入元素0
	heights = append([]int{0}, heights...)
	// 数组尾部加入元素0
	heights = append(heights, 0)
	res := 0
	// 第一个元素已经入栈，从下标1开始
	for i := 1; i < len(heights); i++ {
		// 注意 heights[i] 是和 heights[st[len(st)-1]] 比较 ，st[len(st)-1]] 是下标
		for heights[i] < heights[st[len(st)-1]] {
			mid := st[len(st)-1]
			st = st[:len(st)-1]
			left := st[len(st)-1]
			right := i
			w := right - left - 1
			h := heights[mid]
			res = utils.Max(res, w*h)
		}
		st = append(st, i)
	}
	return res
}

func main() {
	largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
}
