package main

import (
	"amun/cfg"
	"amun/handlers"
	"flag"
	"log"
	"net/http"
	"time"
)

var (
	addr        = flag.String("listen-address", ":9000", "The address to listen on for HTTP requests.")
	cfgFilePath = flag.String("file-path", "config.yaml", "The absolute file for the configuration file.")
)

func main() {
	flag.Parse()

	y := cfg.ReadConfig(*cfgFilePath)

	mux := http.NewServeMux()

	for _, resp := range y.Responses {
		mux.HandleFunc(resp.Path, handlers.AmunHandler(&resp))
	}

	srv := &http.Server{
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      mux,
		Addr:         *addr,
	}

	log.Printf("Starting server listening on port %s", *addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", *addr, err)
	}
}
