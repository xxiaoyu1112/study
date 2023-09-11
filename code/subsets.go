package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{}
	sub := []int{}
	var dfs func(start int)
	dfs = func(start int){
		res = append(res,append([]int{},sub...))
		for i:=start;i<len(nums);i++{
			sub = append(sub,nums[i])
			dfs(i+1)
			sub = sub[:len(sub)-1]
		}
	}
	dfs(0)
	return res
}

func test(){
	var mapTest = map[string]int{
		"12345-6777":4,
		"2":5,
		"3":6,
	}
	for k:= range mapTest{
		//if v == 4{
		//	fmt.Println(k[:5])
		//}
		fmt.Println(k)
	}
}

func main() {
	//nums := []int{1,2,3}
	//res := subsets(nums)
	//fmt.Println("res:",res)
	test()
}
