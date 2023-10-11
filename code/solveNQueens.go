package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	res := [][]string{}
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	cols := map[int]bool{}
	diag1 := map[int]bool{}
	diag2 := map[int]bool{}

	var dfs func(row int)
	dfs = func(row int){
		if row == n{
			temp := make([]string, len(bd))
			for i := 0; i < n; i++ {
				temp[i] = strings.Join(bd[i], "")
			}
			res = append(res,temp)
			return
		}
		for c:=0;c<n;c++{
			if !cols[c] && !diag1[row+c] && !diag2[row-c]{
				bd[row][c] = "Q"
				cols[c] = true
				diag1[row+c] = true
				diag2[row-c] = true
				dfs(row+1)
				bd[row][c] = "."
				cols[c] = false
				diag1[row+c] = false
				diag2[row-c] = false
			}
		}
	}
	dfs(0)
	return res
}

func main() {
	res := solveNQueens(4)
	fmt.Println("res:",res)
}
