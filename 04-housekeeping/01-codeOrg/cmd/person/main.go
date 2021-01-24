package main

import (
	"fmt"

	arch "github.com/disharjayanth/go-explore-interface-error-cxt-arch/04-housekeeping/01-codeOrg"
	"github.com/disharjayanth/go-explore-interface-error-cxt-arch/04-housekeeping/01-codeOrg/storage/mongo"
	"github.com/disharjayanth/go-explore-interface-error-cxt-arch/04-housekeeping/01-codeOrg/storage/postgresql"
)

func main() {
	dbm := mongo.Mongo{}
	dbpg := postgresql.Postg{}

	p1 := arch.Person{
		First: "James",
		Last:  "Bond",
		Age:   28,
	}

	p2 := arch.Person{
		First: "Miss",
		Last:  "MoneyPenny",
		Age:   25,
	}

	fmt.Println("Saving 2 Person in mongo dbm.....")
	arch.Put(dbm, 1, p1)
	arch.Put(dbm, 2, p2)

	fmt.Println("Acessing mongo map dbm with key 1:", arch.Get(dbm, 1))
	fmt.Println("Acessing mongo map dbm with key 2:", arch.Get(dbm, 2))

	fmt.Println("Saving 2 Person in postg dbpg....")
	arch.Put(dbpg, 1, p1)
	arch.Put(dbpg, 2, p2)

	fmt.Println("Accessing postg map dbpg with key 1:", arch.Get(dbpg, 1))
	fmt.Println("Accessing postg map dbpg with key 2:", arch.Get(dbpg, 2))

	ps := arch.NewPersonService(dbm)

	fmt.Println(ps.GetService(1))
	fmt.Println(ps.GetService(2))
}
