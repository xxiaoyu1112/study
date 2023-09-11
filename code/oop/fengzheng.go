package main

// 通过结构体实现封装

type Animal struct {
	name string
}

func NewAnimal() *Animal {
	return &Animal{}
}

func (p *Animal) SetName(name string) {
	p.name = name
}

func (p *Animal) GetName() string {
	return p.name
}
