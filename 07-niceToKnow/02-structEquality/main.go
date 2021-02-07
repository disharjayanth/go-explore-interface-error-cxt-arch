package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type person struct {
	Name  string
	Age   int
	Hobby *Hobby
}

type Hobby struct {
	Name string
}

func main() {
	p1 := person{
		Name: "Smith",
		Age:  22,
		Hobby: &Hobby{
			Name: "Sing",
		},
	}

	p2 := person{
		Name: "John",
		Age:  24,
		Hobby: &Hobby{
			Name: "Sing",
		},
	}

	p3 := person{
		Name: "Smith",
		Age:  22,
		Hobby: &Hobby{
			Name: "Sing",
		},
	}

	p4 := person{
		Name: "John",
		Age:  24,
		Hobby: &Hobby{
			Name: "Sing",
		},
	}

	fmt.Println("cmp pkg on p1 and p2:", cmp.Equal(p1, p2))
	fmt.Println("cmp pkg on p1 and p3:", cmp.Equal(p1, p3))
	fmt.Println("cmp pkg on p2 and p4:", cmp.Equal(p2, p4))
	fmt.Println("cmp pkg on p3 and p4:", cmp.Equal(p3, p4))
}
