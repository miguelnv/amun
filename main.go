package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr        = flag.String("listen-address", ":9000", "The address to listen on for HTTP requests.")
	cfgFilePath = flag.String("file-path", "config.yaml", "The absolute file for the configuration file.")
)

func main() {
	flag.Parse()

	y := ReadConfig(*cfgFilePath)

	mux := http.NewServeMux()

	for _, resp := range y.Responses {
		mux.HandleFunc(resp.Path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", resp.ContentType)

			if resp.ContainsHeaders(&r.Header) && resp.ContainsParams(r.URL.Query()) {
				if _, err := w.Write([]byte(resp.Template)); err != nil {
					log.Printf("Error while serving response %v", resp.Path)
				}
			} else {
				http.NotFound(w, r)
			}

		})
	}

	log.Printf("Starting server listening on port %s", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", *addr, err)
	}
}
