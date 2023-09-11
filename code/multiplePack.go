package main

import (
	"code/code/utils"
	"log"
)

// 多重背包可以化解为 01 背包
func multiplePack(weight, value, nums []int, bagWeight int) int {

	// 将多重背包摊开就是 01 背包
	for i := 0; i < len(nums); i++ {
		for nums[i] > 1 {
			weight = append(weight, weight[i])
			value = append(value, value[i])
			nums[i]--
		}
	}
	log.Println(weight)
	log.Println(value)

	res := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			res[j] = utils.Max(res[j], res[j-weight[i]]+value[i])
		}
		log.Println(res)
	}

	return res[bagWeight]
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	nums := []int{2, 3, 2}
	bagWeight := 10
	multiplePack(weight, value, nums, bagWeight)
}
