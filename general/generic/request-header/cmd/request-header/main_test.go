package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWhoAmIHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/whoami", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	req.RemoteAddr = "127.0.0.1:1234"
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,es;q=0.8")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(WhoAmIHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var got WhoAmI
	err = json.NewDecoder(rr.Body).Decode(&got)

	if err != nil {
		t.Fatalf("Could not decode json response: %v", err)
	}

	expected := WhoAmI{
		Software:  "Mozilla/5.0 (X11; Linux x86_64)",
		Language:  "en-US,en;q=0.9,es;q=0.8",
		IpAddress: "127.0.0.1",
	}

	if got != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", got, expected)
	}
}
