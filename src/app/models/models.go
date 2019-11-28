package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Gender string

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
