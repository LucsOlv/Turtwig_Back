package config

import "time"

type Config struct {
	JWTSecret        string
	JWTExpiresIn    time.Duration
	ServerPort      string
}

func NewConfig() *Config {
	return &Config{
		JWTSecret:     "your-secret-key", // Em produção, use variáveis de ambiente
		JWTExpiresIn:  24 * time.Hour,
		ServerPort:    ":8080",
	}
}
