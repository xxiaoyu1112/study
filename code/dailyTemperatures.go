package main

import "fmt"

func dailyTemperatures(T []int) []int {
	result := make([]int,len(T))
	stack := []int{}

	for i := len(T)-1;i >= 0 ;i-- {
		for len(stack) != 0 && T[i] >= T[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			result[i] = stack[len(stack)-1] - i
		}
		stack = append(stack,i)
	}
	return result
}

// 单调递减栈
func dailyTemperatures1(num []int) []int {
	ans := make([]int, len(num))
	stack := []int{}
	for i, v := range num {
		// 栈不空，且当前遍历元素 v 破坏了栈的单调性
		for len(stack) != 0 && v > num[stack[len(stack)-1]] {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	t := []int{73,74,75,71,69,72,76,73}
	//res := dailyTemperatures(t)
	res1 := dailyTemperatures1(t)
	fmt.Println("res1:",res1)
}
