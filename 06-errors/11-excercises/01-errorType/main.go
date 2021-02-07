package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) Error() string {
	return "Error from person type"
}

func someFunc(p error) {
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", p.Error())
	fmt.Println(p)
}

func main() {
	p1 := &person{
		name: "Smith",
		age:  22,
	}

	someFunc(p1)
}
