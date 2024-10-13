package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	APPID           string
	APPSecret       string
	APPToken        string
	TableID         string
	UserAccessToken string
	OutputTableID   string
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) Parse() {
	cfg.APPID = os.Getenv("APP_ID")
	cfg.APPSecret = os.Getenv("APP_SECRET")
	cfg.APPToken = os.Getenv("APP_TOKEN")
	cfg.TableID = os.Getenv("TABLE_ID")
	cfg.UserAccessToken = os.Getenv("USER_ACCESS_TOKEN")
	cfg.OutputTableID = os.Getenv("OUTPUT_TABLE_ID")

	if cfg.APPID == "" || cfg.APPSecret == "" || cfg.APPToken == "" || cfg.TableID == "" || cfg.UserAccessToken == "" || cfg.OutputTableID == "" {
		fmt.Println("Error: APP ID and APP Secret are required")
		os.Exit(1)
	}
}
