package utils

func Max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func Min(a, b int) int {
	if b > a {
		return a
	}
	return b
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
