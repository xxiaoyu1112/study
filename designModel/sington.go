package designModel

import "sync"

type single struct {
}

var singleInstance *single

var once sync.Once

func getInstance() *single {
	once.Do(func() {
		singleInstance = &single{}
	})
	return singleInstance
}

func init() {
	getInstance()
}
