package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

type App struct {
	Config *Config
	DB     *gorm.DB
}

func (db *DBConfig) dsn() string {
	return fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=true&loc=Local", db.Username, db.Password, db.Host, db.Database, db.Charset)
}

// NewApp generates a newInstance instance of the core application container from the provided configuration
func NewApp(c *Config) *App {
	db, err := gorm.Open(c.DB.Dialect, c.DB.dsn())
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	return &App{
		Config: c,
		DB:     db,
	}
}
