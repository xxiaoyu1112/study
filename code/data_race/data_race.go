package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// 正确用法
func waitGroup() {
	var wg sync.WaitGroup
	var total int32
	sum := 0
	//wg.Add(10)   // 也可以在for循环之前调用 wg.Add(10)
	for i := 1; i <= 10; i++ {
		wg.Add(1) // for循环内 调用wg.Add(1)
		sum += i
		go func(i int) {
			defer wg.Done() // 每个goroutine完成时,使用 wg.Done() 减少 wg 的内部计数器
			//total += i // 这样会触发竞态，多个 goroutine 同时访问同一块内存及 total 所处的内存
			atomic.AddInt32(&total, int32(i))
			//fmt.Println("total", total) // 这样会触发竞态，多个 goroutine 同时访问同一块内存及 total 所处的内存
		}(i)
	}
	wg.Wait()
	fmt.Printf("total:%d sum: %d\n", total, sum)
}

// main 函数先退出了，开启的 goroutine 根本没有机会执行。所以，常见的解决办法是在最后加一个 Sleep
func test1() {
	var total int64
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func(i int) {
			atomic.AddInt64(&total, int64(i))
		}(i)
	}
	fmt.Printf("total:%d sum: %d\n", total, sum)
}

// i 错误的使用
// 主要的原因是 golang 中允许启动的协程中引用外部的变量。主协程的循环很快就跑完了，而各个协程才开始跑,此时结果大部分都是11 ，也有可能是个别的其他值，说明某些协程在主协程循环跑完前执行了
func test2() {
	for i := 1; i <= 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1e9)
}

// 这里的输出顺序是乱的,但是输出 1 到 10 的值
func test3() {
	cpuNum := runtime.NumCPU()
	fmt.Printf("当前系统CPU核心数量：%d\n", cpuNum)
	goroutineNum := runtime.NumGoroutine()
	fmt.Printf("当前程序中的goroutine数量：%d\n", goroutineNum)
	runtime.GOMAXPROCS(1)
	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1e9)
}

func main() {
	waitGroup()
	fmt.Println("---------------")
	var lock sync.Mutex
	total := 0
	for i := 1; i <= 10; i++ {
		go func() {
			lock.Lock()
			total += i
			lock.Unlock()
		}()
	}
	fmt.Println("total:", total)
	fmt.Println("---------------")
	test1()
	fmt.Println("---------------")
	test2()
	fmt.Println("---------------")
	test3()
}
