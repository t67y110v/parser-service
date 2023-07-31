package server

type Config struct {
	BindAddr            string `toml:"bind_addr"`
	LogLevel            string `toml:"log_level"`
	PostgresDatabaseURL string `toml:"postgres_database_url"`
	MongoDatabaseURL    string `toml:"mongo_database_url"`
}

// Default config values
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
