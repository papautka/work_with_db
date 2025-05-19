package postgres

import (
	"github.com/lpernett/godotenv"
	"log"
	"os"
	"strconv"
)

type DbConfig struct {
	Port     int
	User     string
	Password string
	DB       string
	Host     string
}

type Config struct {
	Dsn DbConfig
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return nil
	}
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return &Config{
		Dsn: DbConfig{
			Port:     port,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DB:       os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
		},
	}
}
