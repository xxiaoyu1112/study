package main

import "fmt"

func pow(x float64, n int) float64 {
	ans := 1.0
	xPow := x
	for n > 0 {
		if n%2 == 1 {
			ans *= xPow
		}
		xPow *= xPow
		n /= 2
	}
	return ans
}

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return pow(x, n)
	} else {
		return 1.0 / pow(x, -n)
	}
}

func main() {
	res := myPow(2.0000, -2)
	fmt.Println(res)
}
