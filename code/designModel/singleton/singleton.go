package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Name string
}

var instance *Singleton

var once sync.Once

func getInstance(str string) *Singleton {
	once.Do(func() {
		instance = &Singleton{str}
	})
	return instance
}

func main() {
	fmt.Println(getInstance("a"))
}
