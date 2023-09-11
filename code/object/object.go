package main

import (
	"fmt"
	"reflect"
)

type P struct {
	name string
}

// 不是用指针定义的方法，是值传递的
func (p P) setname(n string) string {
	p.name = n
	return n
}

func (p P) getname() string {
	return p.name
}

func main() {
	// 指针对象，不是用指针定义的方法，是值传递的
	//p := &P{name: "zhangsan"}
	//fmt.Println(p.setname("lisi"))
	//fmt.Println(p.getname())

	// --------- map 实现 Set -------
	//s := SetFactory()
	//s.Add("beijing")
	//s.Add("shanghai")
	//s.Add("beijing")
	//fmt.Println(s)

	// ---
	//runtime.GC()

	//const (
	//	_ = iota
	//	a = 1 << iota
	//	b
	//)
	//const (
	//	c = iota
	//	d
	//)
	//fmt.Println(a, b, c, d)

	//cnt := 0
	//for i := 1; i <= 100; i++ {
	//	//fmt.Println(i & 1)
	//	if i&1 == 0 {
	//		//fmt.Println(i)
	//		cnt += i
	//	} else {
	//		cnt -= i
	//	}
	//}
	//fmt.Println(cnt)

	//deferfunc()
	//res := deferfunc2(1, 2)
	//fmt.Println(res)
	//var i interface{} = Student{}
	//s := i.(Student)
	//fmt.Println(s.name)

	//wg := sync.WaitGroup{}
	//wg.Add(3)
	//for i := 0; i < 3; i++ {
	//	go func() {
	//		fmt.Println(i)
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()

	//chan1 := make(chan struct{}, 1)
	//go func() {
	//	chan1 <- struct{}{}
	//	chan1 <- struct{}{}
	//}()
	//<-chan1
	student := Student{}
	teacher := Teacher{}
	flag := reflect.DeepEqual(student, teacher)
	fmt.Println(flag)
}

type Student struct {
	name string
	Teacher
}

type Teacher struct {
	name string
}

func deferfunc2(a, b int) (c int) {
	defer func() {
		c++
	}()
	c++
	return a + b
}

func deferfunc() {
	defer func() {
		fmt.Println("a")
	}()
	defer func() {
		fmt.Println("b")
	}()
	panic("c")
	defer func() {
		fmt.Println("d")
	}()
}

type Empty struct {
}
type Set struct {
	m map[string]Empty
}

//产生set的工厂
func SetFactory() *Set {
	return &Set{
		m: map[string]Empty{},
	}
}

//添加元素
func (s *Set) Add(val string) {
	s.m[val] = Empty{}
}
func (s *Set) Remove(val string) {
	delete(s.m, val)
}
func (s *Set) Len() int {
	return len(s.m)
}

//清空set
func (s *Set) Clear() {
	s.m = make(map[string]Empty)
}
