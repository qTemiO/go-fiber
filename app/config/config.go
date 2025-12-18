package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type TomlConfig struct {
	SomeValue       string
	SomeSecondValue string
}

type Config struct {
	Port *string
	Host *string

	TomlConfig *TomlConfig
}

var cfg *Config

func LoadConfig(envPath string, tomlPath string) *Config {
	if envPath == "" {
		envPath = "../.env.production"
	}

	if tomlPath == "" {
		tomlPath = "config/config.toml"
	}

	if cfg != nil {
		return cfg
	}

	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("Не удалось загрузить .env файл, используются значения по умолчанию")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "13001"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	cfg = &Config{
		Port: &port,
		Host: &host,

		TomlConfig: LoadToml(tomlPath),
	}

	return cfg
}

func LoadToml(tomlPath string) *TomlConfig {
	var tomlConfig TomlConfig
	if _, err := toml.DecodeFile(tomlPath, &tomlConfig); err != nil {
		slog.Error(fmt.Sprintf("Error occured loading config.toml %s", err))
		return nil
	}
	return &tomlConfig
}
