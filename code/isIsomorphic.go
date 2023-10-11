package main

func isIsomorphic(s string, t string) bool {
	map1 := make(map[byte]byte)
	map2 := make(map[byte]byte)
	for i := range s {
		if _, ok := map1[s[i]]; !ok {
			// map1保存 s[i] 到 t[j]的映射
			map1[s[i]] = t[i]
		}
		if _, ok := map2[t[i]]; !ok {
			// map2保存 t[i] 到 s[j]的映射
			map2[t[i]] = s[i]
		}
		// 无法映射，返回 false
		if map1[s[i]] != t[i] || map2[t[i]] != s[i] {
			return false
		}
	}
	return true
}

func main() {
	isIsomorphic("foo", "bar")
}
