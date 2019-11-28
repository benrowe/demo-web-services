package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type Gender string

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

type Department struct {
	gorm.Model
	Name      string
	Manager   Employee
	Employees []Employee // from & to dates
}

type Salary struct {
	gorm.Model
	Salary   int
	FromDate *time.Time
	ToDate   *time.Time
}

type Manager struct {
	gorm.Model
	Employee Employee
}

type Title struct {
	gorm.Model
}

// Migrate the database to match the defined model structure
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Employee{}, &Department{}, &Salary{}, &Manager{}, &Title{})
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
