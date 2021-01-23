// Concrete Abstract types
package main

import "fmt"

type person struct {
	name string
}

// pointer of person will be of type human
func (p *person) speak() {
	fmt.Println("I'm a person - this is my name: ", p.name)
}

type secretAgent struct {
	person
	ltk bool
}

// pointer of secretAgent will be of type human
func (sa *secretAgent) speak() {
	fmt.Println("I'm a secret agent - this is my name: ", sa.name)
}

type human interface {
	speak()
}

func humanInterfaceFunc(h human) {
	h.speak()
}

func main() {
	// Concrete type
	p1 := &person{
		name: "MoneyPenny",
	}

	// Concrete type
	sa1 := &secretAgent{
		person: person{
			name: "James Bond",
		},
		ltk: true,
	}

	fmt.Printf("%T is the type and value is: %v \n", p1, p1)
	fmt.Printf("%T is the type and value is: %v \n", sa1, sa1)

	// x, y are type of human which is also abstract type
	var x, y human
	x = p1
	y = sa1
	x.speak()
	y.speak()
	fmt.Println("=========== invoking methods from func which takes in human type as arg ==========")
	humanInterfaceFunc(x)
	humanInterfaceFunc(y)
	fmt.Println("======= concrete type is person or/ secretAgent . Abstract type is human since both type as method speak =====")
	humanInterfaceFunc(p1)
	humanInterfaceFunc(sa1)
}
