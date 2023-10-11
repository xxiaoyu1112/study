package main

import (
	"fmt"
	"time"
)

var POOL = 100

func goroutine1(p chan int) {
	for i := 1; i <= POOL; i++ {
		p <- i
		if i%2 == 1 {
			fmt.Println("goroutine-1:", i)
		}
	}
}

func goroutine2(p chan int) {
	for i := 1; i <= POOL; i++ {
		<-p
		if i%2 == 0 {
			fmt.Println("goroutine-2:", i)
		}
	}
}

func goroutine12() {
	msg := make(chan int)

	go goroutine1(msg)
	go goroutine2(msg)

	time.Sleep(time.Millisecond * 1)

}
