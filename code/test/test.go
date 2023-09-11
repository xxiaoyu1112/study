package main

import (
	"code/code/utils"
	"fmt"
	"math"
)

func Dijkstra() {
	var m, s, d int
	fmt.Scan(&m, &s, &d)

	// 初始化邻接矩阵，表示地图之间的连接关系和距离
	graph := make([][]int, m)
	for i := 0; i < m; i++ {
		graph[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&graph[i][j])
		}
	}

	// 使用 Dijkstra 算法计算最短路径
	dist := make([]int, m)     // 存储最短路径
	visited := make([]bool, m) // 存储已访问的节点
	for i := 0; i < m; i++ {
		dist[i] = math.MaxInt32 // 初始化距离为最大值
	}
	dist[s] = 0 // 起点到起点距离为0

	for i := 0; i < m-1; i++ { // 遍历m-1次，每次确定一个节点的最短路径
		u := -1 // u表示当前距离起点最近的未访问节点
		for j := 0; j < m; j++ {
			if !visited[j] && (u == -1 || dist[j] < dist[u]) {
				u = j
			}
		}
		visited[u] = true

		// 更新起点到所有未访问节点的距离
		for v := 0; v < m; v++ {
			if !visited[v] && graph[u][v] > 0 && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	// 输出起点到终点的最短路径
	fmt.Println(dist[d])
}

func beibaowenti() {
	price := []int{5, 1, 4, 11, 15, 67, 10, 21, 35, 14}        // 打造装备所需金币数
	power := []int{100, 80, 70, 60, 200, 300, 59, 41, 170, 90} // 装备所得战力值

	dp := make([][]int, 11)
	for i := 0; i < 11; i++ {
		dp[i] = make([]int, 101)
	}

	for i := 1; i <= 10; i++ { // 遍历每个物品
		for j := 1; j <= 100; j++ { // 遍历每个容量
			if j >= price[i-1] { // 背包容量足够装下当前物品
				dp[i][j] = utils.Max(dp[i-1][j], dp[i-1][j-price[i-1]]+power[i-1])
			} else { // 背包容量不足装下当前物品
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Println(dp[10][100]) // 输出可以获得的最高战力值
}

func main() {
	beibaowenti()
}
