package postgres

import (
	dbinterface "github.com/disharjayanth/go-explore-interface-error-cxt-arch/04-housekeeping/03-exercises/02-dbinterface"
)

// Postgresdb is a dummy database
type Postgresdb map[int]dbinterface.Person

// Save methods saves person give id as n
func (pdb Postgresdb) Save(n int, p dbinterface.Person) {
	pdb[n] = p
}

// Get returns person with given id as n
func (pdb Postgresdb) Get(n int) dbinterface.Person {
	return pdb[n]
}
