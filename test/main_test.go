package test

import (
	"amun/cfg"
	"amun/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStubHandler(t *testing.T) {
	resp := &cfg.Response{
		ContentType: "application/json",
		Path:        "/hello",
		RawTemplate: []byte(`{"alive": true}`)}

	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.AmunHandler(resp))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
