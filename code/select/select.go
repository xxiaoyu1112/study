package main

import (
	"container/heap"
	"fmt"
	"time"
)

func main() {
	// select 的执行顺序是随机的！！！

	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(time.Millisecond)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(time.Millisecond)
		ch2 <- 2
	}()
	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	case <-time.After(2 * time.Millisecond):
		fmt.Println("timeout")
	}

	// 通过嵌套实现优先级的select
	//priority()

	// 优先队列实现
	//priorityQueue()
}

func priority() {
	highPriority := make(chan string)
	lowPriority := make(chan string)

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			lowPriority <- "low priority message"
		}
	}()

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			highPriority <- "high priority message"
		}
	}()

	for {
		select {
		case msg := <-highPriority:
			fmt.Println(msg)
		default:
			select {
			case msg := <-lowPriority:
				fmt.Println(msg)
			case <-time.After(100 * time.Millisecond):
				fmt.Println("timeout")
			}
		}
	}
}

type Item struct {
	value    string
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func priorityQueue() {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	highPriority := make(chan string)
	lowPriority := make(chan string)

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			lowPriority <- "low priority message"
		}
	}()

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			highPriority <- "high priority message"
		}
	}()

	heap.Push(&pq, &Item{value: "highPriority", priority: 2})
	heap.Push(&pq, &Item{value: "lowPriority", priority: 1})

	for {
		select {
		case msg := <-highPriority:
			if pq[0].value == "highPriority" {
				fmt.Println(msg)
			}
		case msg := <-lowPriority:
			if pq[0].value == "lowPriority" {
				fmt.Println(msg)
			}
		case <-time.After(100 * time.Millisecond):
			fmt.Println("timeout")
		}
	}
}
