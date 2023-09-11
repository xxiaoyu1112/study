package main

import "fmt"

func trailingZeroes(n int) int {
	var x = 0
	for n != 0{
		n /= 5
		x += n
	}
	return x
}

func main() {
	n := 100
	x := trailingZeroes(n)
	fmt.Println("x:",x)
}
