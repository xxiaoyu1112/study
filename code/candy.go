package main

import "fmt"

func candy(ratings []int) int {
	// 分发糖果， 贪心算法
	need := make([]int, len(ratings))
	sum := 0
	// 保证条件1，每人先分一颗糖
	for i := range need {
		need[i] = 1
	}
	// 满足条件2，分高的比相邻的多， 左中右，所以分向后遍历及向前遍历两种
	// 先由前向后遍历，遇到右边大于左边的rating的，糖果+1 -->>
	for i:=0;i<len(ratings)-1;i++ {
		if ratings[i] < ratings[i+1] {
			need[i+1] = need[i] + 1
		}
	}
	// 由后向前遍历，遇到左边大于右边的rating，此时应该考虑左边的应取(need[i-1], need[i]+1)的大值，
	// 如果都已经满足左边分高，且糖果多，就不用另外分配糖果了，节省糖果
	for i:=len(ratings)-1;i>0;i-- {
		if ratings[i-1] > ratings[i] {
			need[i-1] = func(a,b int) int {
				if a > b {return a}
				return b
			}(need[i-1], need[i]+1)
		}
	}
	// 最后统计糖果的总数
	for _, v := range need {
		sum += v
	}
	return sum
}

func main() {
	res := candy([]int{1,2,2})
	fmt.Println("res:",res)
}
