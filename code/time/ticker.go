package main

import (
	"fmt"
	"time"
)

func main11() {
	//创建定时器，每隔1秒后，定时器就会给channel发送一个事件(当前时间)
	ticker := time.NewTicker(time.Second * 1)

	i := 0
	go func() {
		for { //循环
			<-ticker.C
			i++
			fmt.Println("i = ", i)

			//if i == 5 {
			//	ticker.Stop() //停止定时器
			//}
		}
	}() //别忘了()

	//死循环，特地不让main goroutine结束
	select {}
}

func main() {
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	go func() {
		for {
			<-ticker.C
			i++
			fmt.Println("i = ", i)
		}
	}()
	//ticker.Stop()
	select {}
}
