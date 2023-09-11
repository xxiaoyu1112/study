package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})
	o := strings.Join(ss, "")
	c := o[0]
	if c == '0' {
		return "0"
	}
	return o
}

func main() {
	nums := []int{3,30,34,5,9}
	res := largestNumber(nums)
	fmt.Println("res:",res)
}
