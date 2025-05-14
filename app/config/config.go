package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port *string
	Host *string
}

var cfg *Config

func LoadConfig() *Config {
	if cfg != nil {
		return cfg
	}

	err := godotenv.Load("../.env.production")
	if err != nil {
		log.Println("Не удалось загрузить .env файл, используются значения по умолчанию")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8104"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	cfg = &Config{
		Port: &port,
		Host: &host,
	}

	return cfg
}
