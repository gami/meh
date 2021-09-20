package config

import (
	"os"
)

// Database represents configuration about database connection.
type Database struct {
	Host     string
	User     string
	Password string
	Port     int
	Name     string
}

func loadDatabaseConfig(conf *Config) {
	db := &conf.DB
	if conf.AppEnv == envTest {
		db.Host = os.Getenv("TEST_DB_HOST")
		db.User = os.Getenv("TEST_DB_USERNAME")
		db.Password = os.Getenv("TEST_DB_PASSWORD")
		db.Port = 3307
		db.Name = os.Getenv("TEST_DB_NAME")
	} else {
		db.Host = os.Getenv("DB_HOST")
		db.User = os.Getenv("DB_USERNAME")
		db.Password = os.Getenv("DB_PASSWORD")
		db.Port = 3306
		db.Name = os.Getenv("DB_NAME")
	}
}
