package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
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

func NewServer(router http.Handler) *Server {
	srv := &http.Server{
		Handler:      router,
		Addr:         cfg.SvcAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return &Server{srv}
}
