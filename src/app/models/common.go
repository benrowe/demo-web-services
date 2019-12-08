package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

// Migrate the database to match the defined model structure
func Migrate(db *gorm.DB) {
	log.Println("migrating db")
	db.AutoMigrate(&Employee{}, &Department{}, &Salary{}, &Manager{}, &Title{})
}
