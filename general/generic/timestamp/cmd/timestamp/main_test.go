package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTimestampHandler(t *testing.T) {
	mux := AppMux()

	tests := []struct {
		route            string
		expectedStatus   int
		expectedResponse any
	}{
		{
			route:            "/api/2016-12-25",
			expectedStatus:   http.StatusOK,
			expectedResponse: Timestamp{Unix: 1482624000000, UTC: "Sun, 25 Dec 2016 00:00:00 GMT"},
		},
		{
			route:            "/api/1451001600000",
			expectedStatus:   http.StatusOK,
			expectedResponse: Timestamp{Unix: 1451001600000, UTC: "Fri, 25 Dec 2015 00:00:00 GMT"},
		},
		{
			route:            "/api/05%20October%202011",
			expectedStatus:   http.StatusOK,
			expectedResponse: Timestamp{Unix: 1317772800000, UTC: "Wed, 05 Oct 2011 00:00:00 GMT"},
		},
		{
			route:            "/api/05%20October%202011%2C%20GMT",
			expectedStatus:   http.StatusOK,
			expectedResponse: Timestamp{Unix: 1317772800000, UTC: "Wed, 05 Oct 2011 00:00:00 GMT"},
		},
		{
			route:            "/api/this-is-not-a-date",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: TimestampError{Error: "Invalid Date"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.route, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.route, nil)
			if err != nil {
				t.Fatalf("Could not created request: %v", err)
			}

			rr := httptest.NewRecorder()

			mux.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}

			switch tt.expectedResponse.(type) {
			case Timestamp:
				var got Timestamp

				err = json.NewDecoder(rr.Body).Decode(&got)
				if err != nil {
					t.Fatalf("Could not decode json response: %v", err)
				}

				if got != tt.expectedResponse {
					t.Errorf("Handler body no contain: %v got %v", tt.expectedResponse, got)
				}

			case TimestampError:
				var got TimestampError

				err = json.NewDecoder(rr.Body).Decode(&got)
				if err != nil {
					t.Fatalf("Could not decode json response: %v", err)
				}

				if got != tt.expectedResponse {
					t.Errorf("Handler body no contain: %v got %v", tt.expectedResponse, got)
				}

			}

		})
	}
}
