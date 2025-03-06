package config

import "time"

type Config struct {
	Port         string        `env:"PORT" envDefault:"8080"`
	JWTSecret    string        `env:"JWT_SECRET" envDefault:"your-secret-key"`
	JWTExpiresIn time.Duration `env:"JWT_EXPIRES_IN" envDefault:"15m"`
}

func NewConfig() *Config {
	return &Config{
		Port:         "8080",
		JWTSecret:    "your-secret-key",
		JWTExpiresIn: 15 * time.Minute,
	}
}
