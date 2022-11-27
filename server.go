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

func (s *Server) ServeMetrics(addr string) {
	go func() {
		log.Fatal(http.ListenAndServe(addr, promhttp.Handler()))
	}()
}

func NewServer(router http.Handler, addr string) *Server {
	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return &Server{srv}
}
