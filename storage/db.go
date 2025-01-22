package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `env:"HOST" env-default:"localhost"`
	Port     string `env:"PORT" env-default:"5432"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	DBName   string `env:"DB_NAME"`
	SSLMode  string `env:"SSL_MODE" env-default:"disable"`
}

var DB *gorm.DB

func ConnectDatabase(config *Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	DB = db
	fmt.Println("Database connected!")
}
