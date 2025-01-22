package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Host             string `env:"HOST" env-default:"localhost"`
	Port             string `env:"PORT" env-default:"5432"`
	Username         string `env:"USERNAME"`
	Password         string `env:"PASSWORD"`
	DBName           string `env:"DB_NAME"`
	SSLMode          string `env:"SSL_MODE" env-default:"disable"`
	JWTAccessSecret  string `env:"JWT_ACCESS_SECRET"`
	JWTRefreshSecret string `env:"JWT_REFRESH_SECRET"`
	JWTAccessExp     string `env:"JWT_ACCESS_TOKEN_EXPIRY"`
	JWTRefreshExp    string `env:"JWT_REFRESH_TOKEN_EXPIRY"`
}

var (
	cfg  *Config
	once sync.Once
)

func GetConfig() (*Config, error) {
	var err error
	once.Do(func() {
		// Load the .env file into the system environment
		err = godotenv.Load()
		if err != nil {
			fmt.Println("No .env file found or error loading it:", err)
		}

		// Initialize and populate the configuration struct
		cfg = &Config{}
		err = cleanenv.ReadEnv(cfg)
	})
	return cfg, err
}
