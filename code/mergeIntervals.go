package main

import (
	"code/code/utils"
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { // 没有一点重合
			res = append(res, prev)
			prev = cur
		} else { // 有重合
			prev[1] = utils.Max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func main() {
	intervals :=[][]int{{1,3}, {2,6},{8,10}, {15,18}}
	//intervals :=[][]int{{1,3}, {0,4},{8,10}, {15,18}}
	res := merge(intervals)
	fmt.Println("res:",res)
}
