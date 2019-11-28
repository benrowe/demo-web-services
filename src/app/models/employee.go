package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type Employee struct {
	gorm.Model
	FirstName   string
	LastName    string
	BirthDate   *time.Time
	Gender      Gender
	StartDate   *time.Time
	Salaries    []Salary
	Departments []Department
	Titles      []Title
}

func (e Employee) Age() int {
	return time.Now().Year() - e.BirthDate.Year()
}

func FindEmployee(db *gorm.DB, id string) (*Employee, error) {
	pk, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	var model Employee
	err = db.First(&model, Employee{
		Model: gorm.Model{ID: uint(pk)},
	}).Error

	return &model, err
}
