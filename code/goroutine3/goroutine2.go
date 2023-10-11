package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func intStringPrint() {
	var chan1, chan2 = make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			<-chan1
			fmt.Println(1)
			chan2 <- struct{}{}
		}
		<-chan1
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			<-chan2
			fmt.Println("A")
			chan1 <- struct{}{}
		}
	}()
	chan1 <- struct{}{}
	wg.Wait()
}

func oddEvenPrint() {
	var chan1, chan2 = make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i += 2 {
			<-chan1
			println(i)
			chan2 <- struct{}{}
		}
		<-chan1
	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-chan2
			println(i)
			chan1 <- struct{}{}
		}
	}()
	chan1 <- struct{}{}
	wg.Wait()
}

func oneProcess() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func oneProcess2() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func(n int) {
		fmt.Println(n)
		wg.Done()
	}(1)
	go func(n int) {
		fmt.Println(n)
		wg.Done()
	}(2)
	go func(n int) {
		fmt.Println(n)
		wg.Done()
	}(3)
	wg.Wait()
}

func oneProcess3() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	for i := 1; i <= 258; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func oneProcess4() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {

		}
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("脑子进煎鱼了")
}

func main() {
	//oddEvenPrint()
	//oneProcess()
	//oneProcess3()
	oneProcess4()
	n := 1000
	sum := 0
	for i := 0; i < n; i *= 2 {
		for j := 0; j < i; j++ {
			sum++
		}
	}
}

func test() {
	var ch1, ch2 = make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i += 2 {
			<-ch1
			fmt.Println(i)
			ch2 <- struct{}{}
		}
		<-ch1
	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-ch2
			fmt.Println(i)
			ch1 <- struct{}{}
		}
	}()
	ch1 <- struct{}{}
	wg.Wait()
}
