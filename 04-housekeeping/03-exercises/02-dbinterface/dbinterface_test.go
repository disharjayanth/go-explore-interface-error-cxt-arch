package dbinterface

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// You can create a dummy postrges db too, or use gomoc which checks just args
// Mongodb is a dummy database
type Mongodb map[int]Person

// Save methods saves person give id as n
func (mdb Mongodb) Save(n int, p Person) {
	mdb[n] = p
}

// Get returns person with given id as n
func (mdb Mongodb) Get(n int) Person {
	return mdb[n]
}

var person1 Person = Person{
	First: "James",
	Last:  "Bond",
	Age:   28,
	Hobby: []string{"Dancing", "Singing", "Golf", "Driving"},
}

func TestPut(t *testing.T) {
	mongodb := Mongodb{}
	mongodb.Save(1, person1)
}

func TestGet(t *testing.T) {
	mongodb := Mongodb{}

	mongodb.Save(1, person1)
	gotPerson1 := mongodb.Get(1)

	if !cmp.Equal(gotPerson1, person1) {
		t.Errorf("Got %v, Want %v ", gotPerson1, person1)
	}
}
