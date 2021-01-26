package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("This person is speaking now: ", p.first)
}

type human interface {
	speak()
}

func invokeSpeak(h human) {
	fmt.Printf("The type of h : %T\n", h)
	h.speak()
}

func main() {
	p1 := &person{
		first: "Jamie",
	}

	p1.speak()

	fmt.Printf("The type of p1 : %T\n", p1)
	invokeSpeak(p1)
}
