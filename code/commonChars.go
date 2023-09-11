package main

import (
	"code/code/utils"
	"fmt"
	"math"
)

func commonChars(words []string) (ans []string) {
	minFreq := [26]int{}
	for i := range minFreq {
		minFreq[i] = math.MaxInt64
	}
	for _, word := range words {
		freq := [26]int{}
		for _, b := range word {
			freq[b-'a']++
		}
		for i, f := range freq[:] {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return
}

func commonChars1(words []string) []string {
	length := len(words)
	fre := make([][]int, 0) //统计每个字符串的词频
	res := make([]string, 0)
	//统计词频
	for i := 0; i < length; i++ {
		var row [26]int //存放该字符串的词频
		for j := 0; j < len(words[i]); j++ {
			row[words[i][j]-97]++
		}
		temp := row[:]
		fre = append(fre, temp)
	}
	//查找一列的最小值
	for j := 0; j < len(fre[0]); j++ {
		pre := fre[0][j]
		for i := 0; i < len(fre); i++ {
			pre = utils.Min(pre, fre[i][j])
		}
		//将该字符添加到结果集（按照次数）
		tmpString := string(rune(j + 97))
		for i := 0; i < pre; i++ {
			res = append(res, tmpString)
		}
	}
	return res
}

func main() {
	res := commonChars([]string{"bella", "label", "roller"})
	fmt.Println("res:", res)
}
