package main

import "fmt"

func main() {
	m := make(map[int]int32)
	fmt.Println(m)
	m1 := new(map[int]int32)
	fmt.Println(m1)
}
