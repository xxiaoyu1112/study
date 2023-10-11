package main

import (
	"fmt"
	"time"
)

// 判断管道有没有存满
//func main() {
//	// 创建管道
//	output1 := make(chan string, 10)
//	// 子协程写数据
//	go write1(output1)
//	// 取数据
//	for s := range output1 {
//		fmt.Println("res:", s)
//		time.Sleep(time.Second)
//	}
//}

func write1(ch chan string) {
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
}
