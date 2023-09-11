package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type leaderBoard struct {
	Player []*Player
}

func (lb *leaderBoard) Len() int {
	return len(lb.Player)
}

func (lb *leaderBoard) Less(i, j int) bool {
	return lb.Player[i].Score > lb.Player[j].Score
}

func (lb *leaderBoard) Swap(i, j int) {
	lb.Player[i], lb.Player[j] = lb.Player[j], lb.Player[i]
}

func (lb *leaderBoard) addPlayToBoard(name string, score int) {
	player := &Player{
		Name:  name,
		Score: score,
	}
	lb.Player = append(lb.Player, player)
	sort.Sort(lb)
}

func (lb *leaderBoard) getNameByScore(name string) int {
	for _, player := range lb.Player {
		if player.Name == name {
			return player.Score
		}
	}
	return -1
}

func (lb *leaderBoard) getScoreByName(score int) string {
	for _, player := range lb.Player {
		if player.Score == score {
			return player.Name
		}
	}
	return ""
}

func main() {
	lb := &leaderBoard{}
	lb.addPlayToBoard("alice", 90)
	lb.addPlayToBoard("bob", 91)
	fmt.Println(lb.getNameByScore("bob"))
	fmt.Println(lb.getScoreByName(90))
}
