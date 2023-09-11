package main

import "fmt"

// 通过内嵌实现继承

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
}

type Student struct {
	*Person
	Grade int
}

func (s *Student) Learn() {
	fmt.Printf("%s is learning in grade %d.\n", s.Name, s.Grade)
}

// 通过组合实现继承

//type Person struct {
//	Name string
//	Age  int
//}
//
//func (p *Person) SayHello() {
//	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
//}
//
//type Student struct {
//	person Person
//	Grade  int
//}
//
//func (s *Student) Learn() {
//	fmt.Printf("%s is learning in grade %d.\n", s.person.Name, s.Grade)
//}
