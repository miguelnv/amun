package handlers

import (
	"amun/cfg"
	"fmt"
	"log"
	"net/http"
)

func AmunHandler(resp *cfg.Response) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", resp.ContentType)
		if resp.ContainsHeaders(&r.Header) && resp.ContainsParams(r.URL.Query()) {
			if _, err := w.Write(resp.RawTemplate); err != nil {
				log.Printf("Error while serving response %v", resp.Path)
				http.Error(w, fmt.Sprintf("Error while serving response %v", resp.Path), http.StatusUnprocessableEntity)
			}
		} else {
			http.NotFound(w, r)
		}
	}
}
