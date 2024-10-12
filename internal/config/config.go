package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	APPID     string
	APPSecret string
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) Parse() {
	cfg.APPID = os.Getenv("APP_ID")
	cfg.APPSecret = os.Getenv("APP_SECRET")

	if cfg.APPID == "" || cfg.APPSecret == "" {
		fmt.Println("Error: APP ID and APP Secret are required")
		os.Exit(1)
	}
}
