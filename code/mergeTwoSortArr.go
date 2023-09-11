package main

import "fmt"

func mergeTwoSortArr(num1, num2 []int) []int {
	tail := len(num1) - 1
	n2 := len(num2) - 1
	n1 := tail - n2 - 1
	for n1 >= 0 && n2 >= 0 {
		if num1[n1] > num2[n2] {
			num1[tail] = num1[n1]
			n1--
			tail--
		} else {
			num1[tail] = num2[n2]
			n2--
			tail--
		}
	}
	for n1 >= 0 {
		num1[tail] = num1[n1]
		n1--
		tail--
	}
	for n2 >= 0 {
		num1[tail] = num2[n2]
		n2--
		tail--
	}
	return num1
}

func main() {
	fmt.Println(mergeTwoSortArr([]int{1, 2, 3, 0, 0, 0}, []int{4, 5, 6}))
}
