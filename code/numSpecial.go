package main

import "fmt"

func numSpecial(mat [][]int) (ans int) {
	for i, row := range mat {
		cnt1 := 0
		for _, x := range row {
			cnt1 += x
		}
		if i == 0 {
			cnt1--
		}
		if cnt1 > 0 {
			for j, x := range row {
				if x == 1 {
					mat[0][j] += cnt1
				}
			}
		}
	}
	for _, x := range mat[0] {
		if x == 1 {
			ans++
		}
	}
	return
}

func main() {
	mat := [][]int{
		{1,0,0},
		{0,0,1},
		{1,0,0},
	}
	res := numSpecial(mat)
	fmt.Println("res:",res)
}
