package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type Gender string

type Employee struct {
    gorm.Model
    FirstName string
    LastName  string
    BirthDate *time.Time
    Gender Gender
    StartDate *time.Time
    Salaries []Salary
}

type Department struct {
    gorm.Model

}

type Salary struct {
    gorm.Model
    Salary int
    FromDate *time.Time
    ToDate *time.Time
}

type Manager struct {
    gorm.Model
}

type Title struct {
    gorm.Model
}

// Migrate the database to match the defined model structure
func Migrate(db *gorm.DB) {
    db.AutoMigrate(&Employee{}, &Department{}, &Salary{}, &Manager{}, &Title{})

}
