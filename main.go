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

	cfg := handlers.ReadConfig(*cfgFilePath)

	log.Println("mappings from configuration file loaded with success")

	for _, mapping := range cfg.Mappings {
		http.Handle(mapping.Path, http.HandlerFunc(mapping.ConfigHandler))
	}

	// generate handler to dinamycally create add/edit/remove/get mappings
	http.Handle("/mappings", http.HandlerFunc(handlers.AddMappingHandler))

	log.Println("mappings endpoint loaded")

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      nil,
		Addr:         *addr,
	}

	log.Printf("Server listening on port %s", *addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", *addr, err)
	} else if err != nil {
		log.Fatalf("Could not start server on %s: %v\n", *addr, err)
	}
}
