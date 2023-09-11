package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	left := 0
	right := len(arr) - k
	for left < right{
		mid := left + (right - left) / 2
		// 对于当前的mid，对比 mid+k 到 x 的距离
		if x - arr[mid] > arr[mid + k] - x {
			// 左边界可以右移
			left = mid + 1
		}else {
			right = mid
		}
	}
	return arr[left:left+k]
}

func main() {
	arr := []int{1,2,3,4,5}
	res := findClosestElements(arr,4,4)
	fmt.Println("res:",res)
}
