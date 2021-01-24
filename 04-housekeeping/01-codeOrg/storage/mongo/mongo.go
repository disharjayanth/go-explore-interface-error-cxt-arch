package mongo

import arch "github.com/disharjayanth/go-explore-interface-error-cxt-arch/04-housekeeping/01-codeOrg"

// Mongo is a dummy ds for arch.Person with id as int and value as arch.Person
type Mongo map[int]arch.Person

func (m Mongo) Save(n int, p arch.Person) {
	m[n] = p
}

func (m Mongo) Retrieve(n int) arch.Person {
	return m[n]
}
