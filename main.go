package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/miguelnv/amun/handlers"
)

var (
	addr        = flag.String("listen-address", ":9000", "The address to listen on for HTTP requests.")
	cfgFilePath = flag.String("file-path", "config.yaml", "The absolute path for the configuration file.")
)

func main() {
	flag.Parse()

	y := handlers.ReadConfig(*cfgFilePath)

	for _, resp := range y.Responses {
		http.Handle(resp.Path, http.HandlerFunc(resp.ConfigHandler))
	}

	// generate post handler to dinamycally create routes
	// http.Handle("/add")

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      nil,
		Addr:         *addr,
	}

	log.Printf("Starting server listening on port %s", *addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", *addr, err)
	}
}
