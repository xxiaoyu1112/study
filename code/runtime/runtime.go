package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("可用CPU核心数为", runtime.NumCPU())

	// Gosched生成一个处理器，允许其他goroutine先运行。 它不会中止当前的 goroutine，因此当前的 goroutine会自动恢复运行。它的作用就是会把当前协程的优先级降低。
	runtime.Gosched()
}
