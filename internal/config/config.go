package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	APPID     string
	APPSecret string
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) Parse() {
	flag.StringVar(&cfg.APPID, "i", "", "APP ID (required)")
	flag.StringVar(&cfg.APPSecret, "s", "", "APP Secret (required)")
	flag.Parse()

	if cfg.APPID == "" || cfg.APPSecret == "" {
		fmt.Println("Error: APP ID and APP Secret are required")
		flag.Usage()
		os.Exit(1)
	}
}
