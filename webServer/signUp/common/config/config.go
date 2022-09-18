package config

import "os"

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
