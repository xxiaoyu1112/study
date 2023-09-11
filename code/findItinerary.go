package main

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	var (
		//每个起点对应的终点数组
		m  = map[string][]string{}
		res []string
	)

	// 初始化, 把图的邻接表存进字典
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	// 按字典序排序
	for key,value := range m {
		fmt.Println(value)
		sort.Strings(m[key])
	}

	// 从‘JFK’开始深搜，每前进一层就减去一条路径，直到某个起点不存在路径的时候就会跳出 循环进行回溯
	// 相对先找不到路径的一定是放在相对后面
	// 所以当前搜索的起点from会插在当前输出路径的第一个位置
	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			// 这个起点终点已经走完了，或者是个死胡同没有回去的路
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			// 还有路继续走
			tmp := m[curr][0]
			// 去掉这个路
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		// 往下递归到终点之后开始入栈
		// 往上递归返回的入栈就是来时的路反着走回去，这样得到的数组是 路径的倒序
		// 越在链路后面的先入栈
		res = append(res, curr)
	}

	dfs("JFK")
	// 翻转数组
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res) - 1 - i] = res[len(res) - 1 - i], res[i]
	}
	return res
}

func main() {
	var tickets = [][]string{
		{"JFK","SFO"},
		{"JFK","ATL"},
		{"SFO","ATL"},
		{"ATL","JFK"},
		{"ATL","SFO"},
	}
	res := findItinerary(tickets)
	fmt.Println("res:",res)
}
