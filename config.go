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
	DbConn      string
	DbUser      string
	DbPass      string
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
		DbConn:      os.Getenv("DB_CONN"),
		DbUser:      os.Getenv("DB_USER"),
		DbPass:      os.Getenv("DB_PASS"),
		LogLvl:      os.Getenv("LOG_LVL"),
	}
}

func GetConfig() *Config {
	return cfg
}
