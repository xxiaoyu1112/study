package main

import "fmt"

func canFormArray(arr []int, pieces [][]int) bool {
	index := make(map[int]int, len(pieces))
	for i, p := range pieces {
		index[p[0]] = i
	}
	for i := 0; i < len(arr); {
		j, ok := index[arr[i]]
		if !ok {
			return false
		}
		for _, x := range pieces[j] {
			if arr[i] != x {
				return false
			}
			i++
		}
	}
	return true
}

func main() {
	res := canFormArray([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}})
	fmt.Println("res:", res)
}
