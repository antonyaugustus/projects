package main

import "fmt"

func main() {
	//initialize the pointer
	p := new(person)
	fmt.Println(p.talk())
	fmt.Println(p.walk())

	o := new(policeofficer)
	fmt.Println(o.talk())
	fmt.Println(o.walk())
}

type person struct {
	firstname string
	lastname string
}

type policeofficer struct {
	badgeNumber int
	precinct string
}

type behaviours interface {
	talk() string
	walk() string
	run() string
}

func (p *person) talk() string {
	return "Hi there"
}

func (p *person) walk() int {
	return 10
}

// func [param[list]] [interface func name] [interface func return type]
func (o *policeofficer) talk() string {
	return "Have a good day!"
}

func (o *policeofficer) walk() int {
	return 20
}