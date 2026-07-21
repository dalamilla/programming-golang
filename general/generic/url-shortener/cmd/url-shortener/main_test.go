package main

import (
	"bytes"
	"encoding/json"
	"github.com/dalamilla/programming-golang/general/generic/url-shortener/internal/database"
	"github.com/dalamilla/programming-golang/general/generic/url-shortener/internal/handler"
	"github.com/dalamilla/programming-golang/general/generic/url-shortener/internal/repository"
	"github.com/dalamilla/programming-golang/general/generic/url-shortener/internal/types"
	"go.etcd.io/bbolt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func setupTestDB(t *testing.T) (*bbolt.DB, func()) {
	tmpDir, err := os.MkdirTemp("", "test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	dbPath := filepath.Join(tmpDir, "test.db")
	db, err := database.InitDB(dbPath)
	if err != nil {
		t.Fatalf("Failed to open test db: %v", err)
	}

	cleanup := func() {
		db.Close()
		os.RemoveAll(tmpDir)
	}

	return db, cleanup
}

func TestShorterHandler(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	shorterRepo, _ := repository.NewShorterRepository(db)
	shorterHandler := handler.NewShorterHandler(shorterRepo)
	mux := AppMux(shorterHandler)

	tests := []struct {
		route            string
		method           string
		body             types.ShorterPayload
		expectedStatus   int
		expectedResponse any
	}{
		{
			route:            "/api/shorturl",
			method:           "POST",
			body:             types.ShorterPayload{URL: "go.dev"},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: types.ShorterError{Error: "Invalid URL"},
		},
		{
			route:            "/api/shorturl",
			method:           "POST",
			body:             types.ShorterPayload{URL: "htps://go.dev"},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: types.ShorterError{Error: "Invalid URL"},
		},
		{
			route:            "/api/shorturl",
			method:           "POST",
			body:             types.ShorterPayload{URL: "https://go.devsss"},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: types.ShorterError{Error: "Invalid Hostname"},
		},
		{
			route:            "/api/shorturl",
			method:           "POST",
			body:             types.ShorterPayload{URL: "https://go.dev"},
			expectedStatus:   http.StatusOK,
			expectedResponse: types.Shorter{ShortURL: 1, OriginalURL: "https://go.dev"},
		},
		{
			route:            "/api/shorturl/100",
			method:           "GET",
			expectedStatus:   http.StatusNotFound,
			expectedResponse: types.ShorterError{Error: "No short URL found for the given input"},
		},
		{
			route:            "/api/shorturl/hundred",
			method:           "GET",
			expectedStatus:   http.StatusNotFound,
			expectedResponse: types.ShorterError{Error: "Wrong format"},
		},
		{
			route:            "/api/shorturl/1",
			method:           "GET",
			expectedStatus:   http.StatusFound,
			expectedResponse: "https://go.dev",
		},
	}

	for _, tt := range tests {
		t.Run(tt.route, func(t *testing.T) {
			var body io.Reader = nil
			if tt.method == "POST" {
				jsonBytes, _ := json.Marshal(tt.body)
				body = bytes.NewReader(jsonBytes)
			}
			req, err := http.NewRequest(tt.method, tt.route, body)
			if err != nil {
				t.Fatalf("Could not created request: %v", err)
			}

			if tt.method == "POST" {
				req.Header.Set("Content-Type", "application/json")
			}

			rr := httptest.NewRecorder()

			mux.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}

			switch tt.expectedResponse.(type) {
			case types.Shorter:
				var got types.Shorter

				err = json.NewDecoder(rr.Body).Decode(&got)
				if err != nil {
					t.Fatalf("Could not decode json response: %v", err)
				}

				if got != tt.expectedResponse {
					t.Errorf("Handler body no contain: %v got %v", tt.expectedResponse, got)
				}

			case types.ShorterError:
				var got types.ShorterError

				err = json.NewDecoder(rr.Body).Decode(&got)
				if err != nil {
					t.Fatalf("Could not decode json response: %v", err)
				}

				if got != tt.expectedResponse {
					t.Errorf("Handler body no contain: %v got %v", tt.expectedResponse, got)
				}
			case string:
				expected := rr.Header().Get("Location")
				if tt.expectedResponse != expected {
					t.Errorf("Handler returned unexpected body: got %v want %v", tt.expectedResponse, expected)
				}

			}

		})
	}

}
