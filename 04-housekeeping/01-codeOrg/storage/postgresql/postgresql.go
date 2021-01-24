package postgresql

import arch "github.com/disharjayanth/go-explore-interface-error-cxt-arch/04-housekeeping/01-codeOrg"

// Postg is a dummy ds for postgresql with key as int and value of type Person
type Postg map[int]arch.Person

func (pg Postg) Save(n int, p arch.Person) {
	pg[n] = p
}

func (pg Postg) Retrieve(n int) arch.Person {
	return pg[n]
}
