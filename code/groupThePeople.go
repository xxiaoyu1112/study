package main

import "fmt"

func groupThePeople(groupSizes []int) [][]int {
	ans := [][]int{}
	group := map[int][]int{}
	for i,size := range groupSizes{
		group[size] = append(group[size],i)
	}
	for size,people := range group{
		for i:=0;i<len(people);i+=size{
			ans = append(ans,people[i:i+size])
		}
	}
	return ans
}

func main() {
	groupSizes := []int{3,3,3,3,3,1,3}
	res := groupThePeople(groupSizes)
	fmt.Println("res:",res)
}
