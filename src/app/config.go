package app

import (
	"log"
	"strconv"
)

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Charset  string
}

type Config struct {
	DB *DBConfig
}

type Environment interface {
	LookupEnv(key string) (string, bool)
}

func LoadConfigFromEnv(e Environment) *Config {
	dbPort, err := strconv.Atoi(env(e, "DB_PORT", "80"))
	if err != nil {
		log.Fatalf("unable to convert requested db port to integer: %v", err)
	}
	return &Config{DB: &DBConfig{
		Dialect:  env(e, "DB_DIALECT", "mysql"),
		Host:     env(e, "DB_HOST", "127.0.0.1"),
		Port:     dbPort,
		Username: env(e, "DB_USERNAME", "root"),
		Password: env(e, "DB_PASSWORD", "password"),
		Database: env(e, "DB_DATABASE", "database"),
		Charset:  env(e, "DB_CHARSET", "utf8"),
	}}
}

// env attempts to get the environment variable, or defaults back if that variable is NOT set
func env(env Environment, key, defaultValue string) string {
	v, ok := env.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return v
}
