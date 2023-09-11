package main

import "fmt"

func permute(nums []int) [][]int {
	var res, path = make([][]int, 0), make([]int, 0)
	var used = make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			var temp = make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := range nums {
			// 剪枝，判断重复使用的数字
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs()
			// 回溯的过程中，将当前的节点从 path 中删除
			path = path[:len(path) - 1]
			used[i] = false
		}
	}

	dfs()

	return res
}

func main() {
	nums := []int{1,2,3}
	res := permute(nums)
	fmt.Println("res:",res)
}
