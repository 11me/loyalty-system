package main

import (
	"log"
	"loyalty-system/pkg/logger"
)

func main() {
	err := logger.InitLogger(cfg.LogLvl)
	if err != nil {
		log.Fatal("Failed to initialize logger.")
	}
	srv := NewServer()
	srv.ServeMetrics()
	log.Fatal(srv.ListenAndServe())
}
