package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/miguelnv/amun/handlers"
)

func TestSuccessStubHandler(t *testing.T) {
	resp := handlers.Mapping{
		ContentType: "application/json",
		Path:        "/hello",
		Template:    `{"alive": true}`,
	}

	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(http.HandlerFunc(resp.ConfigHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestTimerStubHandler(t *testing.T) {
	resp := handlers.Mapping{
		ContentType: "application/json",
		Path:        "/hello",
		Template:    `{"alive": true}`,
	}

	req, err := http.NewRequest("GET", "/unknown", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(http.HandlerFunc(resp.ConfigHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
