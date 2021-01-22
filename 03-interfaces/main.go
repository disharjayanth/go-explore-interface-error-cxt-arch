package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type About interface {
	Info() string
}

func (p *Person) Info() string {
	return fmt.Sprintf("The name of person is %s and is %d years old.", p.Name, p.Age)
}

func InfoOnSomething(a About) string {
	return a.Info()
}

func main() {
	person := &Person{
		Name: "Smith",
		Age:  20,
	}

	fmt.Println("This is from info method from Person type:", person.Info())
	fmt.Println("This is from InfoOnSomething interface:", InfoOnSomething(person))
}
