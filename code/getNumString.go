package main

import (
	"fmt"
)

func getNum(num int) string {
	if num < 10000 {
		return fmt.Sprintf("%v", num)
	}
	mean := float32(num)/float32(10000)
	str := fmt.Sprintf("%v", mean)
	return str + "万"
}

func main() {
	res := getNum(300)
	fmt.Println("res:",res)
}

