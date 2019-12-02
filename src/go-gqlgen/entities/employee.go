package entities

import "time"

// Employee contains data about a individual
type Employee struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	Age         int    `json:"age"`
	Gender      Gender `json:"gender"`
}

func (e *Employee) CurrentRole() *AssignedRole {
	// todo implement
	ar := AssignedRole{
		Role:      &Role{},
		StartDate: time.Now().Format("2006/01/02"),
	}

	return &ar
}

func (e *Employee) Remuneration() *Remuneration {
	// todo
	r := Remuneration{
		Amount: 5,
	}

	return &r
}

func (e *Employee) Department() *Department {
	// todo
	d := Department{
		Name: "sdf",
		Manager: &Employee{
			FirstName: "Harry",
			LastName:  "Tild",
		},
	}

	return &d
}
