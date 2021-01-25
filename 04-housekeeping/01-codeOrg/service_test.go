package architecture

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

type Mongo map[int]Person

func (m Mongo) Save(n int, p Person) {
	m[n] = p
}

func (m Mongo) Retrieve(n int) Person {
	return m[n]
}

func TestPut(t *testing.T) {
	ctl := gomock.NewController(t)
	acc := NewMockAccessor(ctl)
	person := Person{
		First: "Jamie",
		Last:  "Jones",
		Age:   29,
	}

	acc.EXPECT().Save(2, person).MinTimes(1).MaxTimes(1)

	Put(acc, 2, person)

	ctl.Finish()
}

// ExampleTest(in this case its Example*nameoftest*) doesnt take any args and doesnt return anything
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
	// Output: {John Snow 33}
}
