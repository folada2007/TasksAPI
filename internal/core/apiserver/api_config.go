package apiserver

type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_lvl"`
}

func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "info",
	}
}
