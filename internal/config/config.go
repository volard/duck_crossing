package config

type Config struct {
	ServerPort string
	DBConfig   DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadConfig() (*Config, error) {
	// Configuration loading logic
	return nil, nil
}
