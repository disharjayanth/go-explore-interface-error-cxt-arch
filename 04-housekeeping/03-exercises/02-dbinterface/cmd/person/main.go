package main

import (
	"fmt"

	dbinterface "github.com/disharjayanth/go-explore-interface-error-cxt-arch/04-housekeeping/03-exercises/02-dbinterface"
	"github.com/disharjayanth/go-explore-interface-error-cxt-arch/04-housekeeping/03-exercises/02-dbinterface/storage/mongo"
	"github.com/disharjayanth/go-explore-interface-error-cxt-arch/04-housekeeping/03-exercises/02-dbinterface/storage/postgres"
)

func main() {
	mongodb := mongo.Mongodb{}
	postgresdb := postgres.Postgresdb{}

	person1 := dbinterface.Person{
		First: "James",
		Last:  "Bond",
		Age:   28,
		Hobby: []string{"Dancing", "Singing", "Golf", "Driving"},
	}

	person2 := dbinterface.Person{
		First: "Miss",
		Last:  "MoneyPenny",
		Age:   25,
		Hobby: []string{"Singing"},
	}

	fmt.Println("Saving in mongodb")
	dbinterface.Put(mongodb, 1, person1)
	dbinterface.Put(mongodb, 2, person2)

	fmt.Println("Person1 from mongodb:", dbinterface.Retrieve(mongodb, 1))
	fmt.Println("Person2 from mongodb:", dbinterface.Retrieve(mongodb, 2))

	fmt.Println("Saving in postgres")
	dbinterface.Put(postgresdb, 1, person1)
	dbinterface.Put(postgresdb, 2, person2)

	fmt.Println("Person1 from postgres:", dbinterface.Retrieve(postgresdb, 1))
	fmt.Println("Person2 from postgres:", dbinterface.Retrieve(postgresdb, 2))

	personService := dbinterface.NewPersonService(mongodb)
	fmt.Println("PersonService Retrieve from mongodb person1:", personService.GetService(1))
	fmt.Println("PersonService Retrieve from mongodb person2:", personService.GetService(2))

	person3 := dbinterface.Person{
		First: "Smith",
		Last:  "John",
		Age:   30,
		Hobby: []string{"Play", "cook", "sing", "dance"},
	}

	personService.SaveService(3, person3)
	fmt.Println("Person Service Retr from mongodb person3:", personService.GetService(3))
}
