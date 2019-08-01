package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const latencyHeader = "X-Amun-Latency"

// ConfigHandler - defines handler on yaml configuration
func (resp Response) ConfigHandler(w http.ResponseWriter, r *http.Request) {

	if resp.ContainsHeaders(&r.Header) && resp.ContainsParams(r.URL.Query()) {
		w.Header().Set("Content-Type", resp.ContentType)
		wait(&r.Header)

		if _, err := w.Write(resp.RawTemplate); err != nil {
			log.Printf("Error while serving response %v", resp.Path)
			http.Error(w, fmt.Sprintf("Error while serving response %v", resp.Path), http.StatusUnprocessableEntity)
		}
	} else {
		http.NotFound(w, r)
	}
}

func wait(rHds *http.Header) {
	latencyHeaderValue := rHds.Get(latencyHeader)

	if latencyHeaderValue != "" {
		if latency, err := time.ParseDuration(latencyHeaderValue); err == nil {
			timer := time.NewTimer(latency)
			<-timer.C
		}
	}
}
