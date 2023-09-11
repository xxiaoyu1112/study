package main

import (
	"code/code/utils"
	"fmt"
	"math"
	"strconv"
)

func maxPossibleDifference(num int) int {
	numStr := strconv.Itoa(num)
	numStr1 := strconv.Itoa(num)
	n := len(numStr)
	maxDiff := 0
	if n == 1 {
		x, _ := strconv.Atoi(numStr + "5")
		y, _ := strconv.Atoi("5" + numStr1)
		return utils.Abs(x - y)
	}
	for i := 0; i < n; i++ {
		newNum, _ := strconv.Atoi(numStr[:i] + "5" + numStr[i:])
		newNum2, _ := strconv.Atoi(numStr1[:i+1] + "5" + numStr[i+1:])
		if newNum == 0 || newNum2 == 0 {
			continue
		}
		diff := int(math.Abs(float64(newNum - newNum2)))
		fmt.Println(newNum, newNum2)
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	return maxDiff
}

func getMaxSum(numOnes, numZeros, numNegOnes, k int) int {
	ones := make([]int, numOnes)
	for i := 0; i < numOnes; i++ {
		ones[i] = 1
	}
	zeros := make([]int, numZeros)
	for i := 0; i < numZeros; i++ {
		zeros[i] = 0
	}
	negOnes := make([]int, numNegOnes)
	for i := 0; i < numNegOnes; i++ {
		negOnes[i] = -1
	}
	nums := append(append(ones, zeros...), negOnes...)
	count := 0
	sum := 0
	for i := 0; i < k; i++ {
		if nums[i] != 0 {
			count++
			sum += nums[i]
		}
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		if nums[i-k] != 0 {
			count--
			sum -= nums[i-k]
		}
		if nums[i] != 0 && count < k {
			count++
			sum += nums[i]
		}
		maxSum = utils.Max(maxSum, sum)
	}
	return maxSum
}
func minDeletions(word string) int {
	freq := make(map[rune]int)
	for _, ch := range word {
		freq[ch]++
	}
	uniqueFreq := make(map[int]bool)
	for _, f := range freq {
		uniqueFreq[f] = true
	}
	maxFreq := 0
	for f := range uniqueFreq {
		if f > maxFreq {
			maxFreq = f
		}
	}
	if len(uniqueFreq) == len(freq) {
		// 所有字母出现了唯一的次数
		return 0
	} else {
		// 一些字母出现了多次
		return len(word) - maxFreq
	}
}

func isPrime(p int) bool {
	if p <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(p))); i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

//func decrypt(s string) string {
//	n := len(s)
//	original := make([]byte, n)
//	for i := 0; i < n; i++ {
//		original[i] = s[i] - 3
//		if original[i] < 'a' {
//			original[i] += 26
//		}
//	}
//	return string(original)
//}
//
//func main() {
//	var n int
//	fmt.Scan(&n)
//	var s string
//	fmt.Scan(&s)
//	fmt.Println(decrypt(s))
//}

//func main() {
//	//fmt.Println(maxPossibleDifference(143))
//	//fmt.Println(minDeletions("aabbbcc"))
//	//fmt.Println(getMaxSum(3, 2, 0, 4))
//	var a []int
//	b := []int{}
//	a = append(a, 1)
//	b = append(b, 1)
//}
//func kSort(nums []int, k int) int {
//	n := len(nums)
//	swaps := 0
//	for i := 0; i < n; i++ {
//		j := i + 1
//		for ; j < n && j <= i+k; j++ {
//			if nums[j] < nums[i] {
//				nums[i], nums[j] = nums[j], nums[i]
//				swaps++
//			}
//		}
//		if j == n {
//			sort.Ints(nums[i+1:])
//			break
//		}
//	}
//	return swaps
//}
//
//func main() {
//	var t int
//	fmt.Scan(&t)
//	for i := 0; i < t; i++ {
//		var n, k int
//		fmt.Scan(&n, &k)
//		nums := make([]int, n)
//		for j := 0; j < n; j++ {
//			fmt.Scan(&nums[j])
//		}
//		swaps := kSort(nums, k)
//		fmt.Println(swaps)
//	}
//}

//func main() {
//	var n, m int
//	fmt.Scan(&n)
//	arr := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&arr[i])
//	}
//	fmt.Scan(&m)
//	l := make([]int, m)
//	r := make([]int, m)
//	ops := make([]byte, m)
//	x := make([]int, m)
//
//	for i := 0; i < m; i++ {
//		fmt.Scan(&l[i])
//	}
//	for i := 0; i < m; i++ {
//		fmt.Scan(&r[i])
//	}
//	var opsStr string
//	fmt.Scan(&opsStr)
//	ops = []byte(opsStr)
//	for i := 0; i < m; i++ {
//		fmt.Scan(&x[i])
//	}
//
//	for i := 0; i < m; i++ {
//		switch ops[i] {
//		case '=':
//			for j := l[i] - 1; j < r[i]; j++ {
//				arr[j] = x[i]
//			}
//		case '|':
//			for j := l[i] - 1; j < r[i]; j++ {
//				arr[j] |= x[i]
//			}
//		case '&':
//			for j := l[i] - 1; j < r[i]; j++ {
//				arr[j] &= x[i]
//			}
//		}
//	}
//
//	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), " "), "[]"))
//}
//func main() {
//	// 读取输入
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan()
//	n := 0
//	fmt.Sscanf(scanner.Text(), "%d", &n)
//	strs := make([]string, n)
//	for i := 0; i < n; i++ {
//		scanner.Scan()
//		strs[i] = scanner.Text()
//	}
//
//	// 统计重组字符串的数量
//	count := 0
//	m := make(map[string]int)
//	for _, str := range strs {
//		for i := 0; i < len(str); i++ {
//			key := sortedString(str[:i] + str[i+1:])
//			m[key]++
//		}
//	}
//	for _, v := range m {
//		count += v * (v - 1) / 2
//	}
//
//	// 输出结果
//	fmt.Println(count)
//}

//func sortedString(s string) string {
//	chars := []rune(s)
//	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
//	return string(chars)
//}
//func main() {
//	var n int
//	fmt.Scanln(&n)
//
//	var a = make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&a[i])
//	}
//
//	var b = make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&b[i])
//	}
//
//	var c = make([]int, n)
//	for i := 1; i <= n; i++ {
//		c[i-1] = i
//	}
//
//	for i := 0; i < n; i++ {
//		for j := i + 1; j < n; j++ {
//			if b[i] > b[j] && c[i] <= c[j] {
//				c[i], c[j] = c[j], c[i]
//			}
//		}
//	}
//
//	sum := 0
//	for i := 0; i < n; i++ {
//		sum += abs(a[i] - c[i])
//	}
//
//	fmt.Println(sum)
//}

//const mod = 1e9 + 7
//
//func main() {
//	var s string
//	fmt.Scan(&s)
//	res := strings.Split(s, ".")
//	fmt.Println(res)
//}

//func main() {
//	n := 0
//	fmt.Scan(&n)
//	arr := make([]int, n)
//	for i := range arr {
//		fmt.Scan(&arr[i])
//	}
//	sort.Ints(arr)
//	res := 0
//	for i := 0; i < n-1; i++ {
//		res += abs(arr[i] - arr[i+1])
//	}
//	fmt.Println(res)
//}
//
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	var n int
	fmt.Scan(&n)

	capacity := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&capacity[i])
	}

	water := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&water[i])
	}

	cost := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cost[i])
	}

	var m int
	fmt.Scan(&m)

	for i := 0; i < m; i++ {
		var q int
		fmt.Scan(&q)

		// 计算需要倒多少水
		diff := capacity[q-1] - water[q-1]

		// 如果杯子已经是满的，则不需要消耗法力值
		if diff == 0 {
			fmt.Printf("0 ")
			continue
		}

		// 动态规划，求解最小消耗的法力值
		dp := make([]int, diff+1)
		for j := 1; j <= diff; j++ {
			dp[j] = math.MaxInt32
			for k := 0; k < n; k++ {
				if k != q-1 && capacity[k]-water[k] >= j {
					dp[j] = min(dp[j], dp[j-1]+cost[k])
				}
			}
			dp[j] = min(dp[j], dp[j-1]+cost[q-1])
		}

		fmt.Printf("%d ", dp[diff])
	}

	fmt.Println()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
