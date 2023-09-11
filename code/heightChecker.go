package main

import "fmt"

func heightChecker(heights []int) (ans int) {
	cnt := [101]int{}
	for _, v := range heights {
		cnt[v]++
	}

	idx := 0
	for i, c := range cnt {
		for ; c > 0; c-- {
			if heights[idx] != i {
				ans++
			}
			idx++
		}
	}
	return
}

func main() {
	heights := []int{1,1,4,2,1,3}
	ans := heightChecker(heights)
	fmt.Println("ans:",ans)
}
