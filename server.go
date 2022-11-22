package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"loyalty-system/handler"
	"net/http"
	"time"
)

type Server struct {
	*http.Server
}

func (s *Server) ServeMetrics() {
	go func() {
		log.Fatal(http.ListenAndServe(cfg.MetricsAddr, promhttp.Handler()))
	}()
}

func NewServer() *Server {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Get("/inspect/health", handler.GetHealth())

	srv := &http.Server{
		Handler:      router,
		Addr:         cfg.SvcAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return &Server{srv}
}
