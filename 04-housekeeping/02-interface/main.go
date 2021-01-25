package main

import (
	"fmt"
)

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p *person) speak() {
	fmt.Println("From person :", p.first)
}

func (sa *secretAgent) speak() {
	fmt.Println("This is not my real name because im a secret agent : ", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Smith",
	}

	sa := secretAgent{
		person: person{
			first: "James Bond",
		},
		ltk: true,
	}

	saySomething(&p1)
	saySomething(&sa)
}
