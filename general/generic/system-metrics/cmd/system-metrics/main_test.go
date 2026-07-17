package main

import (
	"encoding/json"
	"github.com/dalamilla/programming-golang/general/generic/system-metrics/internal/metrics"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSystemMetricsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/system", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SystemMetricsHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var got metrics.SystemMetrics
	err = json.NewDecoder(rr.Body).Decode(&got)

	if err != nil {
		t.Fatalf("Could not decode json response: %v", err)
	}

	expected := "linux"

	if got.OS != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", got.OS, expected)
	}
}
