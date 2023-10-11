package main

import (
	"fmt"
	"sync"
)

//func goroutine3() {
//	var ch1, ch2, ch3 = make(chan struct{}), make(chan struct{}), make(chan struct{})
//	var wg sync.WaitGroup
//	wg.Add(3)
//	go func() {
//		defer wg.Done()
//		for i := 1; i <= 10; i++ {
//			<-ch1
//			fmt.Println("A")
//			ch2 <- struct{}{}
//		}
//		<-ch1
//	}()
//	go func() {
//		defer wg.Done()
//		for i := 1; i <= 10; i++ {
//			<-ch2
//			fmt.Println("B")
//			ch3 <- struct{}{}
//		}
//	}()
//	go func() {
//		defer wg.Done()
//		for i := 1; i <= 10; i++ {
//			<-ch3
//			fmt.Println("C")
//			ch1 <- struct{}{}
//		}
//	}()
//	ch1 <- struct{}{}
//	wg.Wait()
//}

func goroutine3() {
	var ch1, ch2, ch3 = make(chan struct{}), make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			<-ch1
			fmt.Println("A")
			ch2 <- struct{}{}
		}
		<-ch1
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			<-ch2
			fmt.Println("B")
			ch3 <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			<-ch3
			fmt.Println("C")
			ch1 <- struct{}{}
		}
	}()
	ch1 <- struct{}{}
	wg.Wait()
}
