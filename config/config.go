package config

import (
	"github.com/joho/godotenv"
	"os"
)

var cfg *Config

type Config struct {
	SvcAddr     string
	MetricsAddr string
	DbDSN       string
	LogLvl      string
	Secret      string
}

func InitConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	cfg = &Config{
		SvcAddr:     os.Getenv("SVC_ADDR"),
		MetricsAddr: os.Getenv("METRICS_ADDR"),
		DbDSN:       os.Getenv("DB_DSN"),
		LogLvl:      os.Getenv("LOG_LVL"),
		Secret:      os.Getenv("SECRET"),
	}
	return nil
}

func GetConfig() *Config {
	return cfg
}
