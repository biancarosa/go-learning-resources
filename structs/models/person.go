package models

type person struct {
	FirstName string
	LastName  string
	FullName  string
	Age       int
}

func NewPerson(firstName, lastName string) person {
	p := person{
		FirstName: firstName,
		LastName:  lastName,
	}

	p.FullName = p.FirstName + " " + p.LastName

	return p
}
