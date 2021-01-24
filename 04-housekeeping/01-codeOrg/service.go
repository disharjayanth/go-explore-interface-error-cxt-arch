package architecture

import (
	"fmt"
)

// Person with first, last as string fields and Age int
type Person struct {
	First string
	Last  string
	Age   int
}

type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}

func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}

// PersonService with access field as type of Accessor
type PersonService struct {
	Access Accessor
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{
		Access: a,
	}
}

func (ps PersonService) GetService(n int) (Person, error) {
	p := ps.Access.Retrieve(n)
	if p.First == "" {
		return p, fmt.Errorf("The Person with given id %d is not present in dbm ", n)
	}
	return p, nil
}
