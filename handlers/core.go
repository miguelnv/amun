package handlers

import (
	"github.com/miguelnv/amun/cfg"
	"fmt"
	"log"
	"net/http"
	"time"
)

func CoreHandler(resp cfg.Response) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		timer := getTimer(&r.Header)

		w.Header().Set("Content-Type", resp.ContentType)
		if resp.ContainsHeaders(&r.Header) && resp.ContainsParams(r.URL.Query()) {
			if timer != nil {
				<-timer.C
			}

			if _, err := w.Write(resp.RawTemplate); err != nil {
				log.Printf("Error while serving response %v", resp.Path)
				http.Error(w, fmt.Sprintf("Error while serving response %v", resp.Path), http.StatusUnprocessableEntity)
			}
		} else {
			http.NotFound(w, r)
		}
	}

	return http.HandlerFunc(fn)
}

func getTimer(rHds *http.Header) *time.Timer {
	latency, err := time.ParseDuration(rHds.Get("X-latency"))
	if err == nil {
		return time.NewTimer(latency)
	}

	return nil
}
