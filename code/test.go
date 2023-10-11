package main

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

//func findMaxAverage(nums []int, k int) []float32 {
//	sum := 0
//	for i := 0; i < k; i++ {
//		sum += nums[i]
//	}
//
//	maxSum := sum
//	startIndex := 0
//
//	for i := k; i < len(nums); i++ {
//		sum += nums[i] - nums[i-k]
//		if sum > maxSum {
//			maxSum = sum
//			startIndex = i - k + 1
//		}
//	}
//
//	maxAverage := float32(maxSum) / float32(k)
//
//	result := make([]float32, k)
//	for i := 0; i < k; i++ {
//		result[i] = float32(nums[startIndex+i])
//	}
//
//	//fmt.Printf("子数组的平均值为: %v\n", maxAverage)
//	fmt.Println(maxAverage)
//	return result
//}

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

//func main() {
//	i := GetValue()
//	switch i.(type) {
//	case int:
//		print("int")
//	}
//}
//
//func GetValue() interface{} {
//	return 1
//}

//func main() {
//	fmt.Println(test111(2))
//}
//
//func test111(n int) (res int) {
//	if n >= 5 {
//		return n
//	}
//	return test111(n+1) + test111(n+2)
//}

//func josephus(n, m int) []int {
//	soldiers := make([]int, n)
//	for i := 0; i < n; i++ {
//		soldiers[i] = i + 1
//	}
//
//	result := make([]int, 0, n)
//	index := 0
//
//	for len(soldiers) > 0 {
//		index = (index + m - 1) % len(soldiers)
//		result = append(result, soldiers[index])
//		soldiers = append(soldiers[:index], soldiers[index+1:]...)
//	}
//
//	return result
//}
//
//func main() {
//	var n, m int
//	fmt.Print("请输入士兵的数量n和报数规则m：")
//	fmt.Scanln(&n, &m)
//
//	result := josephus(n, m)
//
//	fmt.Println("士兵出列的顺序为：")
//	for _, soldier := range result {
//		fmt.Printf("%d ", soldier)
//	}
//	fmt.Println()
//}

//func isStraight(cards []int) bool {
//	sort.Ints(cards)
//	gap, zero := 0, 0
//
//	for i := 0; i < len(cards)-1; i++ {
//		if cards[i] == 0 {
//			zero++
//			continue
//		}
//		if cards[i] == cards[i+1] {
//			return false
//		}
//		gap += cards[i+1] - cards[i] - 1
//	}
//
//	return gap <= zero
//}
//
//func main() {
//	nums1 := []int{1, 2, 3, 4, 5}
//	fmt.Println(isStraight(nums1)) // 输出：true
//	nums2 := []int{0, 0, 1, 10, 11}
//	fmt.Println(isStraight(nums2)) // 输出：true
//}
//
//func isPossible(cards []int, points int) bool {
//	if len(cards) == 1 {
//		return cards[0] == points
//	}
//
//	n := len(cards)
//	// 尝试所有可能的运算符组合
//	for i := 0; i < n; i++ {
//		for j := 0; j < n; j++ {
//			if i == j {
//				continue
//			}
//			// 枚举两个不同位置的数字
//			a := cards[i]
//			b := cards[j]
//			// 计算两个数字的运算结果
//			cards[j] = cards[n-1] // 将最后一个数字放到 j 的位置
//			// 加法
//			cards[i] = a + b
//			if isPossible(cards[:n-1], points) {
//				return true
//			}
//			// 减法
//			cards[i] = a - b
//			if isPossible(cards[:n-1], points) {
//				return true
//			}
//			// 乘法
//			cards[i] = a * b
//			if isPossible(cards[:n-1], points) {
//				return true
//			}
//			// 除法（除数不能为0）
//			if b != 0 && a%b == 0 {
//				cards[i] = a / b
//				if isPossible(cards[:n-1], points) {
//					return true
//				}
//			}
//			// 恢复原始值
//			cards[i] = a
//			cards[j] = b
//		}
//	}
//
//	return false
//}
//
//func main() {
//	// 测试例子1
//	cards1 := []int{1, 2, 3, 8}
//	points1 := 24
//	fmt.Println(isPossible(cards1, points1)) // 输出: true
//
//	// 测试例子2
//	cards2 := []int{2, 2, 2, 3, 4}
//	points2 := 24
//	fmt.Println(isPossible(cards2, points2)) // 输出: true
//
//	// 测试例子3
//	cards3 := []int{1, 1, 1, 1}
//	points3 := 5
//	fmt.Println(isPossible(cards3, points3)) // 输出: false
//}

//type echo struct {
//	a string
//}
//
//func main() {
//	var m map[echo]int
//	p := echo{"b"}
//	fmt.Println(m[p])
//}

//func main() {
//	var arr = [3]int{2, 3, 3}
//	var r [3]int
//	for i, v := range arr {
//		if i == 0 {
//			arr[1] = 4
//			arr[2] = 5
//			r[i] = v + 1
//		} else {
//			r[i] = v
//		}
//	}
//	fmt.Print(r)
//}

//func main() {
//	var wg sync.WaitGroup
//	var cnt int
//	wg.Add(1)
//	go func() {
//		time.Sleep(time.Millisecond)
//		cnt++
//		wg.Done()
//		wg.Add(1)
//	}()
//	fmt.Println(cnt)
//	wg.Wait()
//}

//type version struct {
//	i int
//}
//
//func main() {
//	s1 := []int{1, 2, 3}
//	s2 := []int{1, 2, 3}
//	fmt.Println(reflect.DeepEqual(s1, s2))
//	var b1 []byte = nil
//	b2 := []byte{}
//	fmt.Println(reflect.DeepEqual(b1, b2))
//	v1 := version{}
//	v2 := version{}
//
//	fmt.Println(reflect.DeepEqual(v1, v2))
//}

//func main() {
//	//tickets := [][]int{
//	//	{1, 2, 3, 4, 5, 6, 7},
//	//	{4, 6, 7, 9, 10, 12, 14},
//	//}
//	history := [][]int{
//		{1, 2, 3, 4, 5, 6, 7},
//		{8, 9, 10, 11, 12, 13, 14},
//		{15, 16, 17, 18, 19, 20, 21},
//		{22, 23, 24, 25, 26, 27, 28},
//	}
//	fmt.Println(minFrequencyNumbers(history))
//}
//
//func minFrequencyNumbers(tickets [][]int) []int {
//	counts := make(map[int]int, 33)
//	for _, ticket := range tickets {
//		for _, number := range ticket {
//			counts[number]++
//		}
//	}
//
//	type pair struct {
//		number, count int
//	}
//	pairs := make([]pair, 0, len(counts))
//	for number, count := range counts {
//		pairs = append(pairs, pair{number, count})
//	}
//	fmt.Println(pairs)
//	sort.Slice(pairs, func(i, j int) bool {
//		if pairs[i].count == pairs[j].count {
//			return pairs[i].number < pairs[j].number
//		}
//		return pairs[i].count < pairs[j].count
//	})
//	t := findMissingNumbers(counts)
//	result := make([]int, 0, 7)
//	if len(t) >= 7 {
//		result = t[:7]
//	} else {
//		result = t
//		i := 0
//		for len(result) < 7 {
//			result = append(result, pairs[i].number)
//			i++
//		}
//		sort.Ints(result)
//	}
//	return result
//}
//func findMissingNumbers(numbersMap map[int]int) []int {
//	allNumbers := make([]int, 33)
//	for i := 1; i <= 33; i++ {
//		allNumbers[i-1] = i
//	}
//
//	missingNumbers := make([]int, 0)
//	for _, num := range allNumbers {
//		if _, ok := numbersMap[num]; !ok {
//			missingNumbers = append(missingNumbers, num)
//		}
//	}
//	fmt.Println(missingNumbers)
//	return missingNumbers
//}

//// 树的节点
//type Node struct {
//	value    int
//	parent   *Node
//	children []*Node
//}
//
//// 从根节点到当前节点的路径上所有节点的编号
//type Path struct {
//	nodes []*Node
//}
//
//// 创建新的节点
//func NewNode(value int) *Node {
//	return &Node{value: value}
//}
//
//// 创建新的路径
//func NewPath() *Path {
//	return &Path{nodes: make([]*Node, 0)}
//}
//
//// 添加节点到路径
//func (p *Path) AddNode(node *Node) {
//	p.nodes = append(p.nodes, node)
//}
//
//// 计算给定值与路径上所有节点的编号异或的最大值
//func (p *Path) MaxXor(value int) int {
//	maxXor := 0
//	for _, node := range p.nodes {
//		xor := node.value ^ value
//		if xor > maxXor {
//			maxXor = xor
//		}
//	}
//	return maxXor
//}

//func main() {
//	// 读取并解析父节点数组
//	var parentsStr string
//	fmt.Scan(&parentsStr)
//	parents := strings.Split(parentsStr, ",")
//
//	nodes := make([]*Node, len(parents))
//	for i := range nodes {
//		nodes[i] = NewNode(i)
//	}
//
//	for i, p := range parents {
//		parent, _ := strconv.Atoi(p)
//		if parent != -1 {
//			nodes[i].parent = nodes[parent]
//			nodes[parent].children = append(nodes[parent].children, nodes[i])
//		}
//	}
//
//	// 读取并解析问题数组
//	var q int
//	fmt.Scan(&q)
//	questions := make([][]int, q)
//	for i := range questions {
//		var node, value int
//		fmt.Scan(&node, &value)
//		questions[i] = []int{node, value}
//	}
//
//	// 回答问题
//	for _, question := range questions {
//		node, value := question[0], question[1]
//		path := NewPath()
//		current := nodes[node]
//		for current != nil {
//			path.AddNode(current)
//			current = current.parent
//		}
//		fmt.Println(path.MaxXor(value))
//	}
//}
//func main() {
//	count := 0
//
//	for num := 1; num <= 1000; num++ {
//		digitSum := 0
//		temp := num
//
//		// 计算各位数字之和
//		for temp > 0 {
//			digitSum += temp % 10
//			temp /= 10
//		}
//
//		if digitSum%10 == 0 {
//			count++
//		}
//	}
//
//	fmt.Println(count)
//}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func main() {
//	var s, n int
//	fmt.Print("请输入阅饼数量和书籍数量：")
//	fmt.Scanln(&s, &n)
//
//	costs := make([]int, n+1)
//	values := make([]int, n+1)
//	for i := 1; i <= n; i++ {
//		var c, v int
//		fmt.Printf("请输入第 %d 本书的阅饼数和价值：", i)
//		fmt.Scanln(&c, &v)
//		costs[i] = c
//		values[i] = v
//	}
//
//	dp := make([][]int, n+1)
//	for i := 0; i <= n; i++ {
//		dp[i] = make([]int, s+1)
//	}
//
//	for i := 1; i <= n; i++ {
//		for j := 1; j <= s; j++ {
//			if j >= costs[i] {
//				dp[i][j] = max(dp[i-1][j], dp[i-1][j-costs[i]]+values[i])
//			} else {
//				dp[i][j] = dp[i-1][j]
//			}
//		}
//	}
//
//	fmt.Println("兑换书籍的最大价值：", dp[n][s])
//}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//
//func main() {
//	var n, m int
//	fmt.Print("请输入迷宫的行数和列数：")
//	fmt.Scanln(&n, &m)
//
//	maze := make([][]byte, n)
//	for i := 0; i < n; i++ {
//		maze[i] = make([]byte, m)
//		fmt.Printf("请输入第 %d 行的迷宫状态：", i+1)
//		fmt.Scanln(&maze[i])
//	}
//
//	dp := make([][]int, n)
//	for i := 0; i < n; i++ {
//		dp[i] = make([]int, m)
//		for j := 0; j < m; j++ {
//			dp[i][j] = math.MaxInt32
//		}
//	}
//
//	for i := 0; i < n; i++ {
//		for j := 0; j < m; j++ {
//			if maze[i][j] == 's' {
//				dp[i][j] = 0
//			} else if maze[i][j] == '.' {
//				if i > 0 {
//					dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
//				}
//				if j > 0 {
//					dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
//				}
//			}
//		}
//	}
//
//	if dp[n-1][m-1] == math.MaxInt32 {
//		fmt.Println(-1)
//	} else {
//		fmt.Println(dp[n-1][m-1])
//	}
//}

//func main() {
//	i := [3]int{10, 10, 10}
//	j := []int(i[2:3])
//	i[2] = 900
//	fmt.Println(j)
//}

//func leastInterval(tasks []byte, n int) int {
//	freq := make(map[byte]int)
//	maxFreq := 0
//	maxCount := 0
//
//	for _, task := range tasks {
//		freq[task]++
//		if freq[task] > maxFreq {
//			maxFreq = freq[task]
//			maxCount = 1
//		} else if freq[task] == maxFreq {
//			maxCount++
//		}
//	}
//
//	partCount := maxFreq - 1
//	partLength := n - (maxCount - 1)
//	emptySlots := partCount * partLength
//	availableTasks := len(tasks) - maxFreq*maxCount
//	idleSlots := max(0, emptySlots-availableTasks)
//
//	return len(tasks) + idleSlots
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func main() {
//	tasks := []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
//	n := 2
//	result := leastInterval(tasks, n)
//	fmt.Println(result) // Output: 8
//}

//type Point struct {
//	x, y int
//}
//
//var dx = []int{-1, 0, 1, 0}
//var dy = []int{0, 1, 0, -1}
//
//func main() {
//	var m, n int
//	fmt.Scan(&m, &n)
//
//	grid := make([][]int, m)
//	for i := 0; i < m; i++ {
//		grid[i] = make([]int, n)
//		for j := 0; j < n; j++ {
//			fmt.Scan(&grid[i][j])
//		}
//	}
//
//	distance := make([][]int, m)
//	for i := 0; i < m; i++ {
//		distance[i] = make([]int, n)
//		for j := 0; j < n; j++ {
//			distance[i][j] = math.MaxInt32
//		}
//	}
//
//	queue := make([]Point, 0)
//	for i := 0; i < m; i++ {
//		if grid[i][0] == 1 {
//			queue = append(queue, Point{i, 0})
//			distance[i][0] = 0
//		}
//	}
//
//	for len(queue) > 0 {
//		curr := queue[0]
//		queue = queue[1:]
//		for i := 0; i < 4; i++ {
//			newX, newY := curr.x+dx[i], curr.y+dy[i]
//			if newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] == 1 && distance[newX][newY] > distance[curr.x][curr.y]+1 {
//				distance[newX][newY] = distance[curr.x][curr.y] + 1
//				queue = append(queue, Point{newX, newY})
//			}
//		}
//	}
//
//	minDistance := math.MaxInt32
//	for i := 0; i < m; i++ {
//		if grid[i][n-1] == 1 && distance[i][n-1] < minDistance {
//			minDistance = distance[i][n-1]
//		}
//	}
//
//	if minDistance == math.MaxInt32 {
//		fmt.Println(-1)
//	} else {
//		fmt.Println(minDistance)
//	}
//}
//var c = make(chan int)
//var a = 0
//
//func main() {
//	go f()
//	c <- 0
//	print(a)
//}
//
//func f() {
//	a = 1024
//	<-c
//}

//func main() {
//	var n, m, l int
//	fmt.Scan(&n, &m)
//
//	matrix := make([][]int, n)
//	for i := 0; i < n; i++ {
//		matrix[i] = make([]int, m)
//		for j := 0; j < m; j++ {
//			fmt.Scan(&matrix[i][j])
//		}
//	}
//
//	maxValue := 0
//	for i := 0; i < n-1; i++ {
//		for j := 0; j < m-1; j++ {
//			for l = 2; i+l <= n && j+l <= m; l++ {
//				value := matrix[i][j] + matrix[i][j+l-1] + matrix[i+l-1][j] + matrix[i+l-1][j+l-1]
//				if value > maxValue {
//					maxValue = value
//				}
//			}
//		}
//	}
//
//	fmt.Println(maxValue)
//}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

//func letterCombinations(digits string) []string {
//	res := []string{}
//	if len(digits) == 0 {
//		return res
//	}
//	temp := ""
//	var dfs func(start int)
//	dfs = func(start int) {
//		if len(temp) == len(digits) {
//			res = append(res, temp)
//			return
//		}
//		letters := phoneMap[string(digits[start])]
//		for i := 0; i < len(letters); i++ {
//			temp += string(letters[i])
//			dfs(start + 1)
//			temp = temp[:len(temp)-1]
//		}
//	}
//	dfs(0)
//	return res
//}
//
//func main() {
//	fmt.Println(letterCombinations("23"))
//}


//func getDigitSum(num int, base int) int {
//	sum := 0
//	for num > 0 {
//		sum += num % base
//		num /= base
//	}
//	return sum
//}
//
//func isNumberValid(num int) bool {
//	decimalSum := getDigitSum(num, 10)
//	hexSum := getDigitSum(num, 16)
//	octalSum := getDigitSum(num, 8)
//
//	return decimalSum == hexSum && decimalSum == octalSum
//}
//
//func findNumbers(M, N int) []int {
//	numbers := make([]int, 0)
//	count := 0
//
//	for i := M; i <= 99999 && count < N; i++ {
//		if isNumberValid(i) {
//			numbers = append(numbers, i)
//			count++
//		}
//	}
//
//	return numbers
//}
//
//func main() {
//	var M, N int
//	fmt.Print("请输入起始数 M 和返回的个数 N（空格分隔）：")
//	_, err := fmt.Scanf("%d %d", &M, &N)
//	if err != nil {
//		fmt.Println("输入错误：", err)
//		return
//	}
//
//	result := findNumbers(M, N)
//	fmt.Println("符合条件的数：")
//	for _, num := range result {
//		fmt.Println(num)
//	}
//}
