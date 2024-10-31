package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config описывает конфигурацию всего приложения
type Config struct {
	Database struct {
		URL      string
		User     string
		Password string
		Name     string
		Port     int
		Host     string
		SSLMode  string
	}
}

// LoadConfig загружает конфигурацию при помощи Viper
func LoadConfig() *Config {
	var cfg Config

	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	return &cfg
}
