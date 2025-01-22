package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
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
		cfg = &Config{}

		// Read configuration from a .env file
		err = cleanenv.ReadConfig(".env", cfg)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}

		// Validate and populate any remaining env-based values
		err = cleanenv.ReadEnv(cfg)
		if err != nil {
			fmt.Println("Error reading environment variables:", err)
		}
	})
	// fmt.Printf("Loaded Config: %+v\n", cfg)
	return cfg, err
}
