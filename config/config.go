package config

import (
	"fmt"
	"os"
	"sync"
)

const (
	envLocal = "local"
	envTest  = "test"
)

// Config represents configuration root.
type Config struct {
	AppEnv string
	DB     Database
}

var config *Config
var once sync.Once

// GetConfig is a function to get Configuration. This loading process occurs once in boot.
func GetConfig() *Config {
	once.Do(func() {
		cfg, err := NewConfig()
		if err != nil {
			panic(err)
		}
		config = cfg
	})

	return config
}

// NewConfig is a function to init and load Configuration from file or environment variables.
func NewConfig() (*Config, error) {
	conf := &Config{}
	conf.AppEnv = envLocal

	loadFromEnv(conf)
	loadDatabaseConfig(conf)

	return conf, nil
}

func loadFromEnv(conf *Config) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "" {
		conf.AppEnv = appEnv
	}

	fmt.Printf("conf.AppEnv: %v\n", conf.AppEnv)
}
