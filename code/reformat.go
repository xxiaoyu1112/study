package main

import (
	"fmt"
	"unicode"
)

func reformat(s string) string {
	sumDigit := 0
	for _, c := range s {
		if unicode.IsDigit(c) {
			sumDigit++
		}
	}
	sumAlpha := len(s) - sumDigit
	if abs(sumDigit-sumAlpha) > 1 {
		return ""
	}
	flag := sumDigit > sumAlpha
	t := []byte(s)
	for i, j := 0, 1; i < len(t); i += 2 {
		if unicode.IsDigit(rune(t[i])) != flag {
			for unicode.IsDigit(rune(t[j])) != flag {
				j += 2
			}
			t[i], t[j] = t[j], t[i]
		}
	}
	return string(t)
}

//func abs(a int) int{
//	if a>0{
//		return a
//	}
//	return -a
//}

func main() {
	s := "covid2019"
	res := reformat(s)
	fmt.Println("res:", res)
}
