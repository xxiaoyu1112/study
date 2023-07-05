package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建一个信号通道，用于接收退出信号
	sigChan := make(chan os.Signal, 1)
	//signal.Ignore(syscall.SIGINT)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	fmt.Println("阻塞进程中...")

	// 使用 select 语句监听信号通道和阻塞通道
	select {
	case <-sigChan:
		fmt.Println("接收到退出信号")
		//case <-make(chan struct{}):
		// 阻塞进程
	}

	fmt.Println("进程已退出")
}

func quit() {
	// 创建一个信号通道，用于接收退出信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 创建一个阻塞通道，用于阻塞当前进程
	blockChan := make(chan struct{})
	// 启动一个协程，用于处理退出信号
	go func() {
		// 等待接收退出信号
		<-sigChan

		// 发送一个空结构体到阻塞通道中，以便解除阻塞
		blockChan <- struct{}{}
	}()

	fmt.Println("阻塞进程中...")

	// 阻塞当前进程，直到接收到退出信号
	<-blockChan

	fmt.Println("进程已退出")
}
