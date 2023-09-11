package main

func wordBreak(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			str := s[j:i]
			if dp[j] && wordMap[str] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	wordBreak(s, wordDict)
}
