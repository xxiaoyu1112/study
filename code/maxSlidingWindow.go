package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	// 存储所有还没有被移除的下标,下标按照从小到大的顺序被存储，并且它们在数组 nums 中对应的值是严格单调递减的
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		// 不断从队首弹出元素，直到队首元素在窗口中为止
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

func main() {
	nums := []int{1,-1}
	res := maxSlidingWindow(nums,1)
	fmt.Println("res:",res)
}
