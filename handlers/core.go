package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

const latencyHeader = "X-Amun-Latency"

// ConfigHandler - defines handler on yaml configuration
func (m Mapping) ConfigHandler(w http.ResponseWriter, r *http.Request) {

	if m.MatchHeaders(&r.Header) && m.MatchParams(r.URL.Query()) {
		w.Header().Set("Content-Type", m.ContentType)
		wait(&r.Header)

		if _, err := w.Write([]byte(m.Template)); err != nil {
			log.Printf("Error while serving response %v", m.Path)
			http.Error(w, fmt.Sprintf("Error while serving response %v", m.Path), http.StatusUnprocessableEntity)
		}
	} else {
		http.NotFound(w, r)
	}
}

func AddMappingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost,
		http.MethodPut:

		m, err := validateMapping(r)
		if err != nil {
			log.Printf("Error adding mapping. Reason: %v", err)
			http.Error(w, "Could not add mapping", http.StatusUnprocessableEntity)
			return
		}

		log.Printf("Adding mapping %v", m)
		http.Handle(m.Path, http.HandlerFunc(m.ConfigHandler))
		break
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func wait(rHds *http.Header) {
	latencyHdrVal := rHds.Get(latencyHeader)

	if latencyHdrVal != "" {
		if latency, err := time.ParseDuration(latencyHdrVal); err == nil {
			timer := time.NewTimer(latency)
			<-timer.C
		} else {
			log.Printf("Error while applying latency %v", latencyHdrVal)
		}
	}
}

func validateMapping(r *http.Request) (Mapping, error) {

	decoder := json.NewDecoder(r.Body)
	var m Mapping
	if err := decoder.Decode(&m); err != nil {
		return m, errors.New("Could not decode mapping")
	}

	if m.Path == "" {
		return m, errors.New("Path can not be empty")
	}

	if m.ContentType == "" {
		return m, errors.New("Content Type can not be empty")
	}

	return m, nil
}
