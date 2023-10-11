package main

import (
	"fmt"
	"strings"
)

func lengthLongestPath(input string) int {
	// 用\n作为分割符将字符串进行分割
	nodes := strings.Split(input, "\n")
	// length 用来存储文件每一层次的字符最大长度
	length := []int{0}
	// res 表示最长路径
	res := 0
	for i := range nodes {
		// t 表示 文件或者目录所在层次，l 表示 当前文件或者目录字符长度
		t, l := getTab(nodes[i])
		// 除去第一个字符 t=0 ，其余所有字符的长度都需要加上父路径的长度
		if t > 0 {
			l += length[t-1]
		}
		// 判断是否是文件，如果是文件，更新最长路径 res 的长度，并且与之前的路径长度做比较，取最大值
		if isFile(nodes[i]) {
			res = max(res, l+t)
		}
		// 判断 当前层次是否比之前最大层次大，如果大则为新目录或者文件，如果小则表示为新的文件路径， 更新当前路径的长度
		if t >= len(length) {
			length = append(length, l)
		} else {
			length[t] = l
		}
	}
	return res
}

func isFile(str string) bool {
	return strings.Contains(str, ".")
}

// 得到带\t开头的子串
func getTab(str string) (t, l int) {
	for i := range str {
		if str[i] == '\t' {
			// \t的个数就代表这个文件或者目录所在的层级
			t++
		} else {
			break
		}
	}
	l = len(str) - t
	return
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {
	input := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	len := lengthLongestPath(input)
	fmt.Println("len:", len)
}
