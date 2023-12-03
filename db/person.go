package db

import (
	"example.com/error-handling/utils"
	"fmt"
)

// Person struct to represent a person with salary and working hours.
type Person struct {
	Name         string
	Salary       float64
	WorkingHours int
}

// AddPerson adds a new person to the MockDb.
func (db *MockDb) AddPerson(p Person) {
	db.persons[p.Name] = p
}

// GetPersons returns all persons in the MockDb.
func (db *MockDb) GetPerson(name string) (p Person, err error) {
	p, ok := db.persons[name]
	if !ok {
		err = fmt.Errorf("query failed: %w", utils.NewDbNotFoundError())
		return
	}

	return
}
