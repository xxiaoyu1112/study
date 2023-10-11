package main

import "fmt"

func subarraysDivByK0(nums []int, k int) int {
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range nums {
		sum += elem
		modulus := (sum % k + k) % k
		ans += record[modulus]
		record[modulus]++
	}
	return ans
}
func subarraysDivByK(A []int, K int) int {
	preSumModK := 0
	count := 0
	hash := map[int]int{0: 1}

	for i := 0; i < len(A); i++ {
		preSumModK = (preSumModK + A[i]) % K
		// 前缀和 为负数 的情况
		if preSumModK < 0 {
			preSumModK += K
		}
		if v, ok := hash[preSumModK]; ok {
			count += v
			hash[preSumModK]++
		} else {
			hash[preSumModK] = 1
		}
	}
	return count
}


func main() {
	nums := []int{4,5,0,-2,-3,1}
	res := subarraysDivByK(nums,5)
	fmt.Println("res:",res)
}
