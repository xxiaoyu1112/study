package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	s := strconv.Itoa(num)
	str := []byte(s)
	l := len(str) - 1
	for i := range str {
		temp := str[i]
		index := i
		for j := l; j >= i+1; j-- {
			if str[j] > temp {
				temp = str[j]
				index = j
			}
		}
		if index != i {
			str[i], str[index] = str[index], str[i]
			break
		}
	}
	n, _ := strconv.Atoi(string(str))
	return n
}

func main() {
	num := 2736
	res := maximumSwap(num)
	fmt.Println("res:",res)
}
