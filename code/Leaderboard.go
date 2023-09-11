package main

import (
	"fmt"
	"sort"
)

// Player 定义一个结构体表示选手
type Player struct {
	name  string
	score int
}

// Leaderboard 定义一个结构体表示排行榜
type Leaderboard struct {
	players []*Player
}

// Len 实现排序接口
func (lb *Leaderboard) Len() int {
	return len(lb.players)
}

func (lb *Leaderboard) Less(i, j int) bool {
	return lb.players[i].score < lb.players[j].score
}

func (lb *Leaderboard) Swap(i, j int) {
	lb.players[i], lb.players[j] = lb.players[j], lb.players[i]
}

// AddPlayer 添加选手到排行榜中
func (lb *Leaderboard) AddPlayer(name string, score int) {
	player := &Player{name, score}
	lb.players = append(lb.players, player)
	sort.Sort(lb)
}

// GetScoreByName 通过名字查询分数
func (lb *Leaderboard) GetScoreByName(name string) int {
	for _, player := range lb.players {
		if player.name == name {
			return player.score
		}
	}
	return -1 // 没有找到该选手
}

// GetNameByScore 通过分数查询名字
func (lb *Leaderboard) GetNameByScore(score int) string {
	for _, player := range lb.players {
		if player.score == score {
			return player.name
		}
	}
	return "" // 没有找到该分数对应的选手
}

func main() {
	lb := &Leaderboard{}
	lb.AddPlayer("Alice", 80)
	lb.AddPlayer("Bob", 90)
	lb.AddPlayer("Charlie", 70)
	for _, v := range lb.players {
		fmt.Println(v)
	}

	fmt.Println(lb.GetScoreByName("Alice")) // 输出 80
	fmt.Println(lb.GetNameByScore(90))      // 输出 Bob
	fmt.Println(lb.GetScoreByName("David")) // 输出 -1
	fmt.Println(lb.GetNameByScore(100))     // 输出 ""
}
