package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var cfg *Config

type Config struct {
	SvcAddr     string
	MetricsAddr string
	DbDSN       string
	LogLvl      string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file. Probably it does not exist.")
	}

	cfg = &Config{
		SvcAddr:     os.Getenv("SVC_ADDR"),
		MetricsAddr: os.Getenv("METRICS_ADDR"),
		DbDSN:       os.Getenv("DB_DSN"),
		LogLvl:      os.Getenv("LOG_LVL"),
	}
}

func GetConfig() *Config {
	return cfg
}
