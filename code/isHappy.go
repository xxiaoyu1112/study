package main

import "fmt"

func isHappy(n int) bool {
	m := map[int]bool{}
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	t := n / 10
	fmt.Println(t)
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func main() {
	getSum(2)
	isHappy(2)
}
