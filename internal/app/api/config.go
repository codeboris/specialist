package api

type Config struct {
	APPPort     string `toml:"app_port"`
	LoggerLevel string `toml:"logger_level"`
}

func NewConfig() *Config {
	return &Config{
		APPPort:     "8000",
		LoggerLevel: "debug",
	}
}
