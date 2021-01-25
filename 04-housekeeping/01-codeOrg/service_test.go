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

	// Compares actual Put func arguments with acc.EXPECT().Save() arguments, if it mismatches then it throws error
	// Input for below .Save() method are input for expected output i.e (2, person)
	acc.EXPECT().Save(2, person).MinTimes(1).MaxTimes(1)

	Put(acc, 2, person)

	// Look at all .EXPECT() and use them
	ctl.Finish()
}

func TestGet(t *testing.T) {
	ctl := gomock.NewController(t)
	acc := NewMockAccessor(ctl)
	person := Person{}

	// Return declares values to be returned by mocked function call
	// Just compares acc.EXPECT().Retrieve(2) to Get(acc, 2) just checks if the args and return types are same
	acc.EXPECT().Retrieve(2).Return(person)

	Get(acc, 2)

	// Look at all .EXPECT() and use them
	ctl.Finish()
}

// ExampleTest(in this case its Example*nameoftest*) doesnt take any args and doesnt return anything
// Dont use Mock with ExampleTest, if incase u need to implement, then create dummy db and use , just
// like below example
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
