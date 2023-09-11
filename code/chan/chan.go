package main

import (
	"fmt"
)

func main() {
	chan1 := make(chan string, 1)
	go func() {
		fmt.Println(<-chan1)
	}()
	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//go func() {
	//	defer wg.Done()
	//	chan1 <- "aah"
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	fmt.Println(<-chan1)
	//}()
	//wg.Wait()

}
