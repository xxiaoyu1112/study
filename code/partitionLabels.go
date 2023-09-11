package main

import "fmt"

func partitionLabels(s string) []int {
	res := []int{}
	marks := [26]int{}
	for i := range s {
		marks[s[i]-'a'] = i
	}
	left, right := 0, 0
	for i := range s {
		if right < marks[s[i]-'a'] {
			right = marks[s[i]-'a']
		}
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
