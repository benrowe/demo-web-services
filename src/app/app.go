package app

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "log"
)

type App struct {
    Config *Config
    DB *gorm.DB
}

func (db *DBConfig) dsn() string {
    return fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=true&loc=Local", db.Username, db.Password, db.Host, db.Database, db.Charset)
}

// NewApp generates a new instance of the core application container from the provided configuration
func NewApp(c *Config) *App  {
    db, err := gorm.Open(c.DB.Dialect, c.DB.dsn())
    if err != nil {
        log.Fatalf("unable to connect to database: %v", err)
    }

    defer db.Close()

    return &App{
        Config: c,
        DB:     db,
    }
}


