package main

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"log"
	"loyalty-system/config"
	"loyalty-system/db/postgres"
	"loyalty-system/handler"
	"loyalty-system/middleware"
	"loyalty-system/pkg/logger"
)

func main() {
	cfg := config.GetConfig()
	if err := logger.InitLogger(cfg.LogLvl); err != nil {
		log.Fatalf("Failed to initialize logger %s.", err.Error())
	}
	l := logger.GetLogger()

	dbConn, err := postgres.NewDB(cfg.DbDSN)
	if err != nil {
		l.Fatalf("Failed to initialize database %s.", err.Error())
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(chiMiddleware.Recoverer)

	router.Get("/inspect/health", handler.GetHealth())
	router.Post("/api/v1/users", handler.PostUser(dbConn))
	router.Post("/api/v1/tokens", handler.AuthUser(dbConn))

	srv := NewServer(router, cfg.SvcAddr)
	srv.ServeMetrics(cfg.MetricsAddr)
	l.Infof("Listening and serving on %s", cfg.SvcAddr)
	l.Infof("Serving metrics on %s", cfg.MetricsAddr)
	log.Fatal(srv.ListenAndServe())
}
