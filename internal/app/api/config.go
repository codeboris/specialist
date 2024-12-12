package api

import "github.com/codeboris/specialist/storage"

type Config struct {
	APPPort     string `toml:"app_port"`
	LoggerLevel string `toml:"logger_level"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		APPPort:     ":8000",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
