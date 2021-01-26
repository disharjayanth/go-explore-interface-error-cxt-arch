package mongo

import (
	dbinterface "github.com/disharjayanth/go-explore-interface-error-cxt-arch/04-housekeeping/03-exercises/02-dbinterface"
)

// Mongodb is a dummy database
type Mongodb map[int]dbinterface.Person

// Save methods saves person give id as n
func (mdb Mongodb) Save(n int, p dbinterface.Person) {
	mdb[n] = p
}

// Get returns person with given id as n
func (mdb Mongodb) Get(n int) dbinterface.Person {
	return mdb[n]
}
