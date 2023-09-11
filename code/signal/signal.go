package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 通过使用信号监听的方式，程序可以在保证任务执行完整性的同时，实现优雅退出和资源释放的逻辑。

//该示例代码中，程序创建了一个 done 通道和一个 quit 通道，当接收到 SIGINT 或 SIGTERM 信号时，程序会触发一个信号处理函数，在该函数中关闭 done 通道，通知主程序优雅退出。
//主程序使用一个 for 循环来执行任务，同时在 select 语句中监听 done 通道。当 done 通道被关闭时，主程序会退出循环，执行完当前的任务后，程序会输出 "Task finished." 并退出。

//SIGINT 和 SIGTERM 都是 Unix 系统中的信号。它们通常用于向进程发送中断信号，要求进程终止。
//SIGINT 是由用户在终端中按下 Ctrl+C 时发送的中断信号，通常用于请求进程终止正在执行的操作。而 SIGTERM 则是由操作系统或其他进程发送的终止信号，通常用于请求进程正常终止并释放资源。
//当进程接收到 SIGINT 或 SIGTERM 信号时，可以通过信号处理函数来执行一些清理工作，并在适当的时候终止进程。
//可以通过 kill 命令向正在运行的进程发送信号， 发送 SIGINT 信号：kill -2 PID ，发送 SIGTERM 信号：kill -15 PID

func main() {
	done := make(chan struct{})
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		fmt.Println("Received signal, shutting down...")
		close(done)
	}()

	for {
		select {
		case <-done:
			fmt.Println("Task finished.")
			return
		default:
			// Execute the task here.
			fmt.Println("Executing task...")
			time.Sleep(time.Second)
		}
	}
}
