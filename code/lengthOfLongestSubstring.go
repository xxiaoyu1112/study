package main

func lengthOfLongestSubstring(s string) int {
	windows := [128]int{}
	n := len(s)
	left, right := 0, 0
	res := 0
	for right < n {
		rightChar := s[right]
		rightCharIndex := windows[rightChar]
		if rightCharIndex > left {
			left = rightCharIndex
		}
		if res < right-left+1 {
			res = right - left + 1
		}
		windows[rightChar] = right + 1
		right++
	}
	return res
}

func main() {
	lengthOfLongestSubstring("abcabcbb")
}
