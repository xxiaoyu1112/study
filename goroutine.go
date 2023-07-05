package main

import (
	"fmt"
	"sync"
)

// waitGroup 帮助 main等待子 Goroutine结束任务
var wg = new(sync.WaitGroup)

// 多少个Goroutine (1 - n)
var n = 3

// 打印数字，从min - max
var min = 1
var max = 100

// 重点是，创建的channel 如何 在两个goroutine之间被访问，以及信号在 channel中的传递
// 这里还使用了main 来进行协调
// first -> 1 -> chan -> 2 -> chan -> 3 -> chan - > main -> first
func main() {
	wg.Add(n)

	// first 作为 main goroutine 与 1号goroutine 之间的桥梁
	first := make(chan int)
	next := first
	for i := 0; i < n; i++ {
		next = Print(next, i)
	}
	first <- min

	// main 用于协调，读取最后一个通道的信号，并传递给first
	for i := range next {
		first <- i
	}

	close(first)
	wg.Wait()
}

// Print pre 表示，number-1 和 number goroutine之间的桥梁
// next return出去，交给下一个goroutine
func Print(pre chan int, number int) (next chan int) {
	next = make(chan int)
	go func(n int) {
		// close 掉channel，for range不再等待
		defer close(next)
		defer wg.Done()
		for i := range pre {
			if i > max {
				return // 退出goroutine
			}
			fmt.Printf("goroutine %d ： %d \n", number+1, i)
			// 自增，传递给下一个 goroutine
			next <- i + 1
		}
	}(number)

	return
}
