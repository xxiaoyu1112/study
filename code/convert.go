package main

import (
	"code/code/utils"
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	// 以 V 字型为一个循环, 循环周期为 n = (2 * numRows - 2) （2倍行数 - 头尾2个）。
	rows := make([]string, numRows)
	n := 2 * numRows - 2
	for i, char := range s {
		// 对于字符串索引值 i，计算 x = i % n 确定在循环周期中的位置。
		x := i % n
		// 则行号 y = min(x, n - x)。
		rows[utils.Min(x, n - x)] += string(char)
	}
	return strings.Join(rows, "")
}

func main() {
	res := convert("PAYPALISHIRING",3)
	fmt.Println("res:",res)
}
