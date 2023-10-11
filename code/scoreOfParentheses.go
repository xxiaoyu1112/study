package main

import (
	"code/code/utils"
	"fmt"
)

func scoreOfParentheses(s string) int {
	st := []int{0}
	for _, c := range s {
		if c == '(' {
			st = append(st, 0)
		} else {
			v := st[len(st)-1]
			st = st[:len(st)-1]
			st[len(st)-1] += utils.Max(2*v, 1)
		}
	}
	return st[0]
}

func main() {
	res := scoreOfParentheses("(()(()))")
	fmt.Println("res:", res)
}
