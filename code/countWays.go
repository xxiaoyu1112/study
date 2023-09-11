package main

import (
	"code/code/utils"
	"sort"
)

func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	var x, mod int64
	mod = 1000000007
	res := [][]int{}
	prev := ranges[0]
	for i := 1; i < len(ranges); i++ {
		cur := ranges[i]
		if cur[0] <= prev[1] {
			prev[1] = utils.Max(cur[1], prev[1])
		} else {
			prev = cur
			res = append(res, prev)
		}
	}
	res = append(res, prev)
	x = 1 << len(res)
	return int(x % mod)
}

func main() {
	//countWays()
}
