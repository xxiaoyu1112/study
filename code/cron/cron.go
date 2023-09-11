package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/robfig/cron"
)

// 任务结构体
type Task struct {
	Name string
	Spec string
	Cmd  func()
}

// 服务器列表
var servers = []string{"127.0.0.1:8080", "127.0.0.1:8081", "127.0.0.1:8082"}

// 添加任务到所有服务器
func AddTaskToAll(t Task) {
	for _, server := range servers {
		AddTask(server, t)
	}
}

// 添加任务到指定服务器
func AddTask(server string, t Task) {
	c := cron.New()
	c.AddFunc(t.Spec, t.Cmd)
	go c.Start() // 启动cron

	// 连接服务器,发送任务信息
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 发送任务名称和cron spec
	fmt.Fprintf(conn, "Name: %s, Spec: %s\n", t.Name, t.Spec)
}

func main() {
	// 定义任务
	t1 := Task{
		Name: "task1",
		Spec: "*/5 * * * * *",
		Cmd:  func() { fmt.Println("Running task1...") },
	}

	// 添加到所有服务器
	AddTaskToAll(t1)

	// 定时输出
	for range time.Tick(time.Second * 10) {
		fmt.Println("Running...")
	}
}
