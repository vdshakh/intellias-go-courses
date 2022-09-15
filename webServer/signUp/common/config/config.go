package config

import "os"

type Config struct {
	DbDriver    string
	DbSource   string
	Endpoint string
	Port   string
}

func NewConfigFromEnv() *Config {
	config := Config{}
	config.LoadEnv()
	return &config
}

func (c *Config) LoadEnv() {
	c.DbDriver = os.Getenv("DB_DRIVER")
	c.DbSource = os.Getenv("DB_SOURCE")
	c.Endpoint = os.Getenv("ENDPOINT")
	c.Port = os.Getenv("PORT")
}
