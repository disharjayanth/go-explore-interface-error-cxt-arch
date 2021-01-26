package dbinterface

// Person struct has four fields: First of type string, Last string, Age int
// Hobby slice of string
type Person struct {
	First string
	Last  string
	Age   int
	Hobby []string
}

// Accessor any type having below methods with same signature is also type of Accessor
type Accessor interface {
	Get(n int) Person
	Save(n int, p Person)
}

// Put takes in Accessor , n int , and p Person to be save with resp n as id
func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}

// Retrieve takes in Accessor, n int and returns a Person with give id as n
func Retrieve(a Accessor, n int) Person {
	return a.Get(n)
}

// PersonService struct as one field of type Accessor, any type having Get and Save
// methods with same signature is also of type Accessor therefore can be in this field
type PersonService struct {
	Access Accessor
}

// NewPersonService creates a new pointer to PersonService
func NewPersonService(a Accessor) *PersonService {
	return &PersonService{
		Access: a,
	}
}

// GetService returns Person with given n
func (ps *PersonService) GetService(n int) Person {
	return ps.Access.Get(n)
}

// SaveService save with person with n as id
func (ps *PersonService) SaveService(n int, p Person) {
	ps.Access.Save(n, p)
}
