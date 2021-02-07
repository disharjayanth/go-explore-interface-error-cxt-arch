package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) String() string {
	return fmt.Sprint(`Name is `, p.name, ` and age is `, p.age)
}

func main() {
	p1 := &person{
		name: "James Bond",
		age:  22,
	}

	fmt.Println(p1)
}
