package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type mongo map[int]person
type postg map[int]person

func (m mongo) save(n int, p person) {
	m[n] = p
}

func (m mongo) retrieve(n int) person {
	return m[n]
}

func (pg postg) save(n int, p person) {
	pg[n] = p
}

func (pg postg) retrieve(n int) person {
	return pg[n]
}

type accessor interface {
	save(n int, p person)
	retrieve(n int) person
}

func put(a accessor, n int, p person) {
	a.save(n, p)
}

func get(a accessor, n int) person {
	return a.retrieve(n)
}

func main() {
	dbm := mongo{}
	dbpg := postg{}

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   28,
	}

	p2 := person{
		first: "Miss",
		last:  "MoneyPenny",
		age:   25,
	}

	fmt.Println("Saving 2 person in mongo dbm.....")
	put(dbm, 1, p1)
	put(dbm, 2, p2)

	fmt.Println("Acessing mongo map dbm with key 1:", get(dbm, 1))
	fmt.Println("Acessing mongo map dbm with key 2:", get(dbm, 2))

	fmt.Println("Saving 2 person in postg dbpg....")
	put(dbpg, 1, p1)
	put(dbpg, 2, p2)

	fmt.Println("Accessing postg map dbpg with key 1:", get(dbpg, 1))
	fmt.Println("Accessing postg map dbpg with key 2:", get(dbpg, 2))
}
