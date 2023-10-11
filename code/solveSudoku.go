package main

import "fmt"

func solveSudoku(board [][]byte)  {
	var dfs func() bool
	dfs = func() bool{
		for i:=0;i<9;i++{
			for j:=0;j<9;j++{
				if board[i][j] != '.'{
					continue
				}
				for k:='1';k<='9';k++{
					if isValid(i,j,byte(k),board) {
						board[i][j] = byte(k)
						if dfs() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	dfs()
}

func isValid(row,col int ,k byte ,board [][]byte) bool{
	for i:=0;i<9;i++{
		if board[row][i] == k || board[i][col] == k{
			return false
		}
	}
	startRow := (row / 3)*3
	startCol := (col / 3)*3
	for i:=startRow;i<startRow+3;i++{
		for j:=startCol;j<startCol+3;j++{
			if board[i][j] == k{
				return false
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}
	solveSudoku(board)
	fmt.Println("board:",board)
}
