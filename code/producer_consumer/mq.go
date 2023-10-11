package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	producerCount     = 3
	consumerCount     = 2
	channelBufferSize = 100
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	dataChannel := make(chan int, channelBufferSize)

	for i := 0; i < producerCount; i++ {
		wg.Add(1)
		go producer(dataChannel, i)
	}

	for i := 0; i < consumerCount; i++ {
		wg.Add(1)
		go consumer(dataChannel, i)
	}

	wg.Wait()
	close(dataChannel)
	fmt.Println("All producers and consumers have finished processing.")
}

func producer(dataChannel chan<- int, id int) {
	defer wg.Done()

	for {
		time.Sleep(1 * time.Second)

		num := rand.Intn(101)
		fmt.Printf("Producer %d produced: %d\n", id, num)

		dataChannel <- num
	}
}

func consumer(dataChannel <-chan int, id int) {
	defer wg.Done()

	for num := range dataChannel {
		fmt.Printf("Consumer %d consumed: %d\n", id, num)
	}
}
