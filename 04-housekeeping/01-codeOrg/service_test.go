package architecture

import (
	"fmt"
	"testing"
)

type Mongo map[int]Person

func (m Mongo) Save(n int, p Person) {
	m[n] = p
}

func (m Mongo) Retrieve(n int) Person {
	return m[n]
}

func TestPut(t *testing.T) {
	mdb := Mongo{}
	person := Person{
		First: "Jamie",
		Last:  "Jones",
		Age:   29,
	}
	Put(mdb, 2, person)

	if got := mdb.Retrieve(2); got != person {
		t.Errorf("Got %v, want %v", got, person)
	}
}

func ExamplePut() {
	mdb := Mongo{}
	person := Person{
		First: "John",
		Last:  "Snow",
		Age:   33,
	}

	Put(mdb, 3, person)
	got := Get(mdb, 3)

	fmt.Println(got)
}
