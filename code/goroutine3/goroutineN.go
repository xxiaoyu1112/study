package main

import "fmt"

func goroutineN() {
	chanNum := 4
	chanQueue := make([]chan struct{}, chanNum)
	var result = 0
	exitChan := make(chan struct{})
	for i := 0; i < chanNum; i++ {
		chanQueue[i] = make(chan struct{})
		if i == chanNum-1 {
			go func(i int) {
				chanQueue[i] <- struct{}{}
			}(i)
		}
	}
	for i := 0; i < chanNum; i++ {
		var lastChan, curChan chan struct{}
		if i == 0 {
			lastChan = chanQueue[chanNum-1]
		} else {
			lastChan = chanQueue[i-1]
		}
		curChan = chanQueue[i]
		go func(i byte, lastChan, curChan chan struct{}) {
			for {
				if result > 20 {
					exitChan <- struct{}{}
				}
				<-lastChan
				fmt.Printf("%c\n", i)
				result++
				curChan <- struct{}{}
			}
		}('A'+byte(i), lastChan, curChan)
	}
	<-exitChan
	fmt.Println("done")
}
