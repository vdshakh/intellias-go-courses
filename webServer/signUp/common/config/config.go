package config

import (
	"errors"
	"os"
)

var (
	errInvalidDBDriver = errors.New("db driver is empty")
	errInvalidDBSource = errors.New("db source is empty")
	errInvalidPort     = errors.New("port is empty")
)

type Config struct {
	DBDriver string
	DBSource string
	Port     string
}

func NewConfigFromEnv() *Config {
	config := Config{}
	config.LoadEnv()

	return &config
}

func (c *Config) LoadEnv() {
	c.DBDriver = os.Getenv("DB_DRIVER")
	c.DBSource = os.Getenv("DB_SOURCE")
	c.Port = os.Getenv("PORT")
}

func (c *Config) Validate() error {
	if c.DBSource == "" {
		return errInvalidDBDriver
	}

	if c.DBDriver == "" {
		return errInvalidDBDriver
	}

	if c.Port == "" {
		return errInvalidPort
	}

	return nil
}
