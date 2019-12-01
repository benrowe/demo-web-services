package entities

// Employee contains data about a individual
type Employee struct {
	ID           string        `json:"id"`
	FirstName    string        `json:"firstName"`
	LastName     string        `json:"lastName"`
	DateOfBirth  string        `json:"dateOfBirth"`
	Age          int           `json:"age"`
	Gender       Gender        `json:"gender"`
	Remuneration *Remuneration `json:"remuneration"`
	Department   *Department   `json:"department"`
}

func (e *Employee) CurrentRole() *AssignedRole {
	// todo implement
	ar := AssignedRole{}

	return &ar
}
