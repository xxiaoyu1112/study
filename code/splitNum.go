package main

import (
	"fmt"
	"sort"
	"strconv"
)

func splitNum(num int) int {
	str := strconv.Itoa(num)
	if len(str) == 2 {
		x, _ := strconv.Atoi(string(str[0]) + string(str[1]))
		return x
	}
	b := []byte(str)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	x, y := "", ""
	for i := 0; i < len(b); i += 2 {
		x += string(b[i])
	}
	fmt.Println(x)
	for j := 1; j < len(b); j += 2 {
		y += string(b[j])
	}
	fmt.Println(y)
	a, _ := strconv.Atoi(x)
	t, _ := strconv.Atoi(y)
	return a + t
}

func main() {
	res := splitNum(10)
	fmt.Println(res)
}
