// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package entities

import (
	"fmt"
	"io"
	"strconv"
)

type AssignRoleInput struct {
	Role      string   `json:"role"`
	Employees []string `json:"employees"`
	StartDate string   `json:"startDate"`
}

// The assignment of a role to an employee
type AssignedRole struct {
	Role      *Role     `json:"role"`
	Employee  *Employee `json:"employee"`
	StartDate string    `json:"startDate"`
	EndDate   *string   `json:"endDate"`
}

type CreateDepartmentInput struct {
	Name    string `json:"name"`
	Manager string `json:"manager"`
}

type CreateEmployeeInput struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	Gender      Gender `json:"Gender"`
}

type CreateRoleInput struct {
	Title string `json:"title"`
}

type Department struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Manager *Employee `json:"manager"`
}

// Employee contains data about a individual
type Employee struct {
	ID           string        `json:"id"`
	FirstName    string        `json:"firstName"`
	LastName     string        `json:"lastName"`
	DateOfBirth  string        `json:"dateOfBirth"`
	Age          int           `json:"age"`
	Gender       Gender        `json:"gender"`
	Remuneration *Remuneration `json:"remuneration"`
	CurrentRole  *AssignedRole `json:"currentRole"`
	Department   *Department   `json:"department"`
}

type Remuneration struct {
	Amount    int              `json:"amount"`
	Type      RenumerationType `json:"type"`
	StartDate string           `json:"startDate"`
	EndDate   *string          `json:"endDate"`
}

type Role struct {
	Title string `json:"title"`
}

type Gender string

const (
	GenderMale      Gender = "MALE"
	GenderFemale    Gender = "FEMALE"
	GenderUndefined Gender = "UNDEFINED"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
	GenderUndefined,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale, GenderUndefined:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RenumerationType string

const (
	RenumerationTypeSalary RenumerationType = "SALARY"
	RenumerationTypeWage   RenumerationType = "WAGE"
)

var AllRenumerationType = []RenumerationType{
	RenumerationTypeSalary,
	RenumerationTypeWage,
}

func (e RenumerationType) IsValid() bool {
	switch e {
	case RenumerationTypeSalary, RenumerationTypeWage:
		return true
	}
	return false
}

func (e RenumerationType) String() string {
	return string(e)
}

func (e *RenumerationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RenumerationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RenumerationType", str)
	}
	return nil
}

func (e RenumerationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
