package main

import (
	"log"
)

func main() {
	srv := NewServer()
	srv.ServeMetrics()
	log.Fatal(srv.ListenAndServe())
}
