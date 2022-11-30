package main

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"loyalty-system/config"
	"loyalty-system/db/postgres"
	"loyalty-system/handler"
	"loyalty-system/middleware"
	"loyalty-system/pkg/logger"
	"net/http"
)

func main() {
	logger.InitLogger()
	log := logger.GetLogger()

	if err := config.InitConfig(); err != nil {
		log.Fatalf("failed to initialize config %s", err.Error())
	}
	cfg := config.GetConfig()

	dbConn, err := postgres.NewDB(cfg.DbDSN)
	if err != nil {
		log.Fatalf("failed to initialize database %s", err.Error())
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(chiMiddleware.Recoverer)

	router.Get("/inspect/health", handler.GetHealth())
	router.Post("/api/v1/users", handler.PostUser(dbConn))
	router.Post("/api/v1/tokens", handler.AuthUser(dbConn))

	router.Mount("/api", protectedRouter(dbConn, cfg))

	srv := NewServer(router, cfg.SvcAddr)
	srv.ServeMetrics(cfg.MetricsAddr)
	log.Infof("Listening and serving on %s", cfg.SvcAddr)
	log.Infof("Serving metrics on %s", cfg.MetricsAddr)
	log.Fatal(srv.ListenAndServe())
}

func protectedRouter(dbConn *postgres.DBConn, cfg *config.Config) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Authenticate(cfg.Secret))
	router.Post("/v1/organizations", handler.PostOrganization(dbConn))

	return router
}
