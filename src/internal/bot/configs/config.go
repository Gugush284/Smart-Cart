package configs

// Config ...
type Config struct {
	LogLevel   string `toml:"log_level"`
	Apitoken   string `toml:"token"`
	ServerAddr string `toml:"server_addr"`
}

// New Config
func NewConfig() *Config {
	return &Config{
		LogLevel:   "debug",
		ServerAddr: "http://localhost:8080",
	}
}
