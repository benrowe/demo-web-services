package models

import "github.com/jinzhu/gorm"

// Migrate the database to match the defined model structure
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Employee{}, &Department{}, &Salary{}, &Manager{}, &Title{})
}
