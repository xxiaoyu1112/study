package main

import (
	"fmt"
)

func minDiff(n, m int, nums [][]int) int {
	total := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			total += nums[i][j]
		}
	}

	dp := make([]bool, total/2+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := total / 2; k >= nums[i][j]; k-- {
				dp[k] = dp[k] || dp[k-nums[i][j]]
			}
		}
	}

	for i := total / 2; ; i-- {
		if dp[i] {
			return total - 2*i
		}
	}
}

//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan()
//	line := scanner.Text()
//	parts := strings.Split(line, " ")
//	n, _ := strconv.Atoi(parts[0])
//	m, _ := strconv.Atoi(parts[1])
//
//	nums := make([][]int, n)
//	for i := 0; i < n; i++ {
//		nums[i] = make([]int, m)
//		scanner.Scan()
//		line = scanner.Text()
//		parts = strings.Split(line, " ")
//		for j := 0; j < m; j++ {
//			nums[i][j], _ = strconv.Atoi(parts[j])
//		}
//	}
//
//	fmt.Println(minDiff(n, m, nums))
//}

//var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右四个方向
//
//func minBlocks(n int, s string) int {
//	var minBlock = n
//	for i := 1; i*i <= n; i++ {
//		if n%i == 0 {
//			minBlock = utils.Min(minBlock, getBlocks(n/i, i, s))
//		}
//	}
//	return minBlock
//}
//
//func getBlocks(row, col int, s string) int {
//	var blocks = 0
//	matrix := make([][]byte, row)
//	visited := make([][]bool, row)
//	for i := 0; i < row; i++ {
//		matrix[i] = []byte(s[i*col : (i+1)*col])
//		visited[i] = make([]bool, col)
//	}
//
//	for i := 0; i < row; i++ {
//		for j := 0; j < col; j++ {
//			if !visited[i][j] {
//				blocks++
//				dfs(i, j, row, col, matrix, visited)
//			}
//		}
//	}
//	return blocks
//}

//func dfs(x, y, row, col int, matrix [][]byte, visited [][]bool) {
//	visited[x][y] = true
//	for _, d := range dir {
//		nx, ny := x+d[0], y+d[1]
//		if nx >= 0 && nx < row && ny >= 0 && ny < col && !visited[nx][ny] && matrix[nx][ny] == matrix[x][y] {
//			dfs(nx, ny, row, col, matrix, visited)
//		}
//	}
//}

//func main() {
//	var n int
//	var s string
//	fmt.Scan(&n)
//	fmt.Scan(&s)
//	s = strings.Trim(s, "\n")
//	fmt.Println(minBlocks(n, s))
//}

//func main() {
//	var n int
//	fmt.Scan(&n)
//
//	weights := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&weights[i])
//	}
//
//	edges := make([][]int, n-1)
//	for i := 0; i < n-1; i++ {
//		var u, v int
//		fmt.Scan(&u, &v)
//		edges[i] = []int{u, v}
//	}
//
//	maxRedNodes := findMaxRedNodes(n, weights, edges)
//	fmt.Println(maxRedNodes)
//}
//
//func findMaxRedNodes(n int, weights []int, edges [][]int) int {
//	graph := make(map[int][]int)
//	for _, edge := range edges {
//		u, v := edge[0], edge[1]
//		graph[u] = append(graph[u], v)
//		graph[v] = append(graph[v], u)
//	}
//
//	maxRedNodes := 0
//	visited := make([]bool, n+1)
//	for i := 1; i <= n; i++ {
//		redNodes := dfs(i, 0, weights, graph, visited)
//		if redNodes > maxRedNodes {
//			maxRedNodes = redNodes
//		}
//	}
//
//	return maxRedNodes
//}
//
//func dfs(node, parent int, weights []int, graph map[int][]int, visited []bool) int {
//	visited[node] = true
//	redNodes := 0
//
//	for _, neighbor := range graph[node] {
//		if neighbor != parent {
//			if isPerfectSquare(weights[node-1] * weights[neighbor-1]) {
//				if !visited[neighbor] {
//					redNodes += dfs(neighbor, node, weights, graph, visited)
//				}
//				redNodes++
//			} else {
//				redNodes += dfs(neighbor, node, weights, graph, visited)
//			}
//		}
//	}
//
//	return redNodes
//}
//
//func isPerfectSquare(num int) bool {
//	sqrt := int(math.Sqrt(float64(num)))
//	return sqrt*sqrt == num
//}

//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
//
//func minStep(n, m, x1, y1, x2, y2, x3, y3 int) int {
//	d1 := min(abs(x1-x2), n-abs(x1-x2)) + min(abs(y1-y2), m-abs(y1-y2))
//	d2 := min(abs(x2-x3), n-abs(x2-x3)) + min(abs(y2-y3), m-abs(y2-y3))
//	return d1 + d2
//}
//
//func main() {
//	var n, m int
//	fmt.Scan(&n, &m)
//	var x1, y1, x2, y2, x3, y3 int
//	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)
//	fmt.Println(minStep(n, m, x1, y1, x2, y2, x3, y3))
//}
//
//var adjacencyList map[int][]int
//var maxCount int
//
//func main() {
//	var n, k int
//	fmt.Scan(&n, &k)
//
//	adjacencyList = make(map[int][]int, n)
//	maxCount = 0
//
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Split(bufio.ScanWords)
//
//	for i := 0; i < n-1; i++ {
//		scanner.Scan()
//		u := parseInt(scanner.Text())
//		scanner.Scan()
//		v := parseInt(scanner.Text())
//
//		adjacencyList[u] = append(adjacencyList[u], v)
//		adjacencyList[v] = append(adjacencyList[v], u)
//	}
//
//	dfs(1, 0, k, -1)
//
//	fmt.Println(maxCount)
//}
//
//func dfs(node, distance, k, parent int) {
//	if distance > k {
//		return
//	}
//
//	count := 1 // Include the current node
//
//	for _, neighbor := range adjacencyList[node] {
//		if neighbor != parent {
//			dfs(neighbor, distance+1, k, node)
//			count += maxCount
//		}
//	}
//
//	if count > maxCount {
//		maxCount = count
//	}
//}
//
//func parseInt(s string) int {
//	var num int
//	var sign int = 1
//	for i := 0; i < len(s); i++ {
//		if s[i] == '-' {
//			sign = -1
//		} else {
//			num = num*10 + int(s[i]-'0')
//		}
//	}
//	return num * sign
//}

//
//func main() {
//	var p float64
//	fmt.Scan(&p)
//
//	// 未触发大保底机制时抽到当期5星的概率
//	prob := p / 2
//
//	// 计算期望次数
//	expect := 0.0
//
//	// 当期5星的概率为 prob，常驻5星的概率为 1 - prob
//	// 若未触发大保底机制，每次抽取有 prob 的概率抽到当期5星，有 1 - prob 的概率抽到常驻5星
//	// 若已触发大保底机制，100%概率抽到当期5星
//	// 若连续89次未抽到5星，则下一抽必定抽到当期5星
//	// 这里采用递归的方式计算期望次数
//	expect = calcExpectation(prob, 89)
//
//	fmt.Printf("%.7f", expect)
//}
//
//func calcExpectation(prob float64, consecutive int) float64 {
//	if consecutive == 0 {
//		return 0
//	}
//
//	// 若连续抽了 consecutive 次未抽到5星，则下一抽必定抽到当期5星
//	if consecutive == 1 {
//		return 1
//	}
//
//	// 若已触发大保底机制，则100%概率抽到当期5星
//	if consecutive <= 89 {
//		return 1 + calcExpectation(prob, consecutive-1)
//	}
//
//	// 若未触发大保底机制，每次抽取有 prob 的概率抽到当期5星，有 1 - prob 的概率抽到常驻5星
//	// 递归计算期望次数
//	return 1/prob + (1-prob)/prob*calcExpectation(prob, consecutive-1)
//}

//func main() {
//	var n int
//	fmt.Scan(&n)
//
//	arr := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&arr[i])
//	}
//
//	maxSum := getMaxSum(arr)
//	fmt.Println(maxSum)
//}
//
//func getMaxSum(arr []int) int {
//	n := len(arr)
//	dp := make([][]int, n)
//	for i := range dp {
//		dp[i] = make([]int, 2)
//	}
//
//	dp[0][0], dp[0][1] = arr[0], arr[0]
//
//	for i := 1; i < n; i++ {
//		dp[i][0] = utils.Max(dp[i-1][0]+arr[i], dp[i-1][1]+arr[i])
//		dp[i][1] = utils.Max(dp[i-1][0]+arr[i-1]*arr[i], dp[i-1][1]+arr[i-1]*arr[i])
//	}
//
//	// Find the maximum sum by considering both cases (with and without magic)
//	maxWithoutMagic := dp[n-1][0]
//	maxWithMagic := dp[n-1][1]
//
//	for i := 1; i < n; i++ {
//		// Try using magic between element i-1 and i
//		magicSum := dp[i-1][0] + arr[i-1]*arr[i] + dp[n-1][1] - dp[i][1]
//		maxWithMagic = utils.Max(maxWithMagic, magicSum)
//	}
//
//	return utils.Max(maxWithoutMagic, maxWithMagic)
//}

//
//func main() {
//	var n int
//	fmt.Scanln(&n)
//
//	heights := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&heights[i])
//	}
//
//	maxDiff := calculateMaxDiff(heights)
//	fmt.Println(maxDiff)
//}
//
//func calculateMaxDiff(heights []int) int {
//	sort.Ints(heights)
//
//	maxDiff := heights[0] - heights[len(heights)-1] // 初始化最大差值为第一个人和最后一个人的身高差值
//
//	for i := 0; i < len(heights)-1; i++ {
//		diff := heights[i+1] - heights[i]
//		if diff > maxDiff {
//			maxDiff = diff
//		}
//	}
//
//	return maxDiff
//}
//
//func abs(n int) int {
//	if n < 0 {
//		return -n
//	}
//	return n
//}

//func findSequence(n int, a []int) []int {
//	b := make([]int, n)
//	used := make(map[int]bool, n)
//
//	for i := 0; i < n; i++ {
//		// 从 1 到 n 中选择一个未使用的值作为 b[i]
//		j := 1
//		for {
//			if !used[j] && (a[i]+j)%(i+1) == 0 {
//				b[i] = j
//				used[j] = true
//				break
//			}
//			x := used[j]
//			if x {
//				j = j + n
//			} else {
//				j = j + 1
//			}
//		}
//	}
//
//	return b
//}
//
//func main() {
//	var n int
//	fmt.Scanln(&n)
//
//	a := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&a[i])
//	}
//
//	b := findSequence(n, a)
//	for i := 0; i < n; i++ {
//		fmt.Printf("%d ", b[i])
//	}
//	fmt.Println()
//}

//func findMaxScore(n, t int, problems [][]int) string {
//	// 创建动态规划数组 dp，dp[i][j] 表示前 i 道题在总时间为 j 时的最大得分
//	dp := make([][]int, n+1)
//	for i := 0; i <= n; i++ {
//		dp[i] = make([]int, t+1)
//	}
//
//	// 动态规划求解最大得分
//	for i := 1; i <= n; i++ {
//		ti1, si1, ti2, si2 := problems[i-1][0], problems[i-1][1], problems[i-1][2], problems[i-1][3]
//		for j := 1; j <= t; j++ {
//			// 暴力算法得分
//			if j >= ti2 {
//				dp[i][j] = dp[i-1][j-ti2] + si2
//			}
//			// 正确算法得分
//			if j >= ti1 && dp[i][j] < dp[i-1][j-ti1]+si1 {
//				dp[i][j] = dp[i-1][j-ti1] + si1
//			}
//		}
//	}
//
//	// 根据动态规划数组构造做题方案
//	strategy := make([]rune, n)
//	j := t
//	for i := n; i > 0; i-- {
//		ti1, s1, ti2, s2 := problems[i-1][0], problems[i-1][1], problems[i-1][2], problems[i-1][3]
//		if j >= ti1 && dp[i][j] == dp[i-1][j-ti1]+s1 {
//			strategy[i-1] = 'A' // 写正确算法
//			j -= ti1
//		} else if j >= ti2 && dp[i][j] == dp[i-1][j-ti2]+s2 {
//			strategy[i-1] = 'B' // 写暴力算法
//			j -= ti2
//		} else {
//			strategy[i-1] = 'F' // 放弃此题
//		}
//	}
//
//	return string(strategy)
//}
//
//func main() {
//	var n, t int
//	fmt.Scan(&n, &t)
//
//	problems := make([][]int, n)
//	for i := 0; i < n; i++ {
//		problems[i] = make([]int, 4)
//		fmt.Scan(&problems[i][0], &problems[i][1], &problems[i][2], &problems[i][3])
//	}
//
//	strategy := findMaxScore(n, t, problems)
//	fmt.Println(strategy)
//}

func findMaxAverage(nums []int, k int) []float32 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum
	startIndex := 0

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
			startIndex = i - k + 1
		}
	}

	maxAverage := float32(maxSum) / float32(k)

	result := make([]float32, k)
	for i := 0; i < k; i++ {
		result[i] = float32(nums[startIndex+i])
	}

	//fmt.Printf("子数组的平均值为: %v\n", maxAverage)
	fmt.Println(maxAverage)
	return result
}

//func main() {
//	nums := []int{1, 2, 3, 1}
//	k := 2
//
//	subarray := findMaxAverage(nums, k)
//	fmt.Printf("子数组: %v\n", subarray)
//}

//var m map[[2]int]int
//
//func main() {
//	m[[2]int{0, 0}] = 1
//	fmt.Println(m)
//}

//func generateString(k int) string {
//	alphabet := "abcdefghijklmnopqrstuvwxyz"
//
//	// 构建一个初始字符串，包含所有字母
//	initialString := alphabet[:k] + alphabet[:k]
//
//	// 构造一个结果字符串，将初始字符串反复添加
//	result := ""
//	for i := 0; i < k; i++ {
//		result += initialString
//	}
//
//	return result
//}
//
//func main() {
//	k := 1
//	result := generateString(k)
//	if result == "" {
//		fmt.Println("无法生成满足条件的字符串")
//	} else {
//		fmt.Println(result)
//	}
//}

//func main() {
//	var n int
//	fmt.Scan(&n)
//
//	monsters := make([]Monster, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&monsters[i].hp)
//		monsters[i].idx = i + 1
//	}
//
//	duels := make([]Duel, 0)
//
//	sort.Slice(monsters, func(i, j int) bool {
//		return monsters[i].hp < monsters[j].hp
//	})
//
//	for len(monsters) > 1 {
//		i := 0
//		if monsters[i].hp == monsters[i+1].hp {
//			fmt.Println("monsters[i].hp == monsters[i+1].hp", monsters)
//			duel := Duel{monsters[i].idx, monsters[i+1].idx}
//			duels = append(duels, duel)
//			monsters = monsters[2:]
//		} else {
//			fmt.Println("monsters:", monsters)
//			duel := Duel{monsters[0].idx, monsters[len(monsters)-1].idx}
//			duels = append(duels, duel)
//			if monsters[0].hp == monsters[len(monsters)-1].hp {
//				monsters = monsters[1 : len(monsters)-1]
//			} else {
//				monsters[len(monsters)-1].hp = abs(monsters[0].hp - monsters[len(monsters)-1].hp)
//				//monsters = monsters[:len(monsters)-1]
//				monsters = monsters[1:]
//				sort.Slice(monsters, func(i, j int) bool {
//					return monsters[i].hp < monsters[j].hp
//				})
//			}
//		}
//	}
//	fmt.Println(monsters)
//	if len(monsters) == 0 {
//		fmt.Println(0)
//	} else {
//		fmt.Println(monsters[0].hp)
//	}
//	fmt.Println(len(duels))
//	for _, duel := range duels {
//		fmt.Println(duel.a, duel.b)
//	}
//}
//
//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
//
//type Monster struct {
//	idx int
//	hp  int
//}
//
//type Duel struct {
//	a, b int
//}
//
//func main() {
//
//	for i := 1; i < 10; i++ {
//		// 输入正整数
//		var inputStr string
//		fmt.Print("请输入一个正整数: ")
//		fmt.Scan(&inputStr)
//
//		// 检查输入是否为正整数
//		num, err := strconv.Atoi(inputStr)
//		if err != nil || num <= 0 {
//			fmt.Println("输入无效，请输入一个正整数。")
//			return
//		}
//
//		// 将输入数字转换为字符串
//		numStr := strconv.Itoa(num)
//
//		// 找到旋转数字的位数
//		maxRotatedNumStr := findMaxRotatedNumber(numStr)
//
//		// 输出结果
//		fmt.Printf("旋转后的最大数为：%s\n", maxRotatedNumStr)
//	}
//}
//
//// 找到旋转数字的位数之后能达到的最大数
//func findMaxRotatedNumber(numStr string) string {
//	// 初始化最大数为输入数本身
//	maxNumStr := numStr
//
//	// 将数字字符串拼接成两倍长度，以便于循环旋转
//	numStr = numStr + numStr
//
//	// 循环旋转数字，比较大小
//	for i := 1; i < len(numStr)/2; i++ {
//		rotatedNumStr := numStr[i : i+len(numStr)/2]
//		if strings.Compare(rotatedNumStr, maxNumStr) > 0 {
//			maxNumStr = rotatedNumStr
//		}
//	}
//
//	return maxNumStr
//}

//func main() {
//	//var n, k int
//	//fmt.Scan(&n, &k)
//	//
//	//var letters string
//	//fmt.Scan(&letters)
//	k := 3
//	n := 2
//	letters := "AB"
//	i, cnt, Acnt := 0, 0, 0
//	for {
//		if Acnt == k {
//			fmt.Println(cnt)
//			return
//		} else {
//			cnt++
//			t := letters[i%n]
//			fmt.Println(string(t))
//			if t == 'A' {
//				Acnt++
//			}
//			i++
//		}
//	}
//}

//const modulo = 1000000007
//
//type Node struct {
//	Value    int
//	Color    int
//	Children []*Node
//}
//
//func main() {
//	var n int
//	fmt.Scan(&n)
//
//	parents := make([]int, n-1)
//	for i := 0; i < n-1; i++ {
//		fmt.Scan(&parents[i])
//	}
//
//	colors := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&colors[i])
//	}
//
//	nodes := make(map[int]*Node)
//	for i := 1; i <= n; i++ {
//		nodes[i] = &Node{Value: 1, Color: colors[i-1]}
//	}
//
//	for i := 2; i <= n; i++ {
//		parent := parents[i-2]
//		nodes[parent].Children = append(nodes[parent].Children, nodes[i])
//	}
//
//	root := nodes[1]
//	result := calculateValue(root)
//	fmt.Println(result)
//}
//
//func calculateValue(node *Node) int {
//	if len(node.Children) == 0 {
//		return 1
//	}
//
//	value := 0
//	for _, child := range node.Children {
//		childValue := calculateValue(child)
//
//		if node.Color == 1 {
//			value = (value + childValue) % modulo
//		} else if node.Color == 2 {
//			value = (value * childValue) % modulo
//		}
//	}
//
//	return value
//}

//func main() {
//	if true {
//		defer fmt.Printf("1")
//	} else {
//		defer fmt.Printf("2")
//	}
//	fmt.Printf("3")
//}

//func main() {
//	strs := []string{"one", "two", "three"}
//	for _, s := range strs {
//		go func() {
//			time.Sleep(1 * time.Second)
//			fmt.Printf("%s", s)
//		}()
//	}
//	time.Sleep(3 * time.Second)
//}

//type TransInfo struct {
//}
//
//type Fragment interface {
//	Exec(transInfo *TransInfo) error
//}
//
//type GetPodAction struct {
//}
//
//func (g GetPodAction) Exec(info *TransInfo) error {
//	return nil
//}
//
//func main() {
//	var fragment1 Fragment = new(GetPodAction)
//	var fragment2 Fragment = &GetPodAction{}
//	var fragment3 Fragment = GetPodAction{}
//	var fragment2 Fragment = GetPodAction
//}

func main() {
	i := GetValue()
	switch i.(type) {
	case int:
		print("int")
	}
}

func GetValue() interface{} {
	return 1
}
