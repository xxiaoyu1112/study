package main

import "fmt"

func main() {
	chan1 := make(chan struct{})
	go func() {
		defer func() {
			recover()
		}()
		chan1 <- struct{}{}
		panic("a")
		fmt.Println("a")
	}()
	select {
	case <-chan1:
		fmt.Println("b")
	}
}
