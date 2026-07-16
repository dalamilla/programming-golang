package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFormHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FormHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	got := rr.Body.String()
	expected := "Form File"
	if !strings.Contains(got, expected) {
		t.Errorf("Handler body no contain: %v got %v", expected, got)
	}
}

func TestFileMetadataHandler(t *testing.T) {
	tests := []FileMetadata{
		{Name: "test.txt", Type: "text/plain", Size: 14},
		{Name: "test.json", Type: "application/json", Size: 40},
		{Name: "test.png", Type: "image/png", Size: 13736},
	}

	for _, expected := range tests {
		t.Run(expected.Name, func(t *testing.T) {
			filePath := filepath.Join("..", "..", "testdata", expected.Name)

			file, err := os.Open(filePath)
			if err != nil {
				t.Fatalf("Failed to open test data file: %v", err)
			}
			defer file.Close()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			filename := filepath.Base(filePath)

			contentType := mime.TypeByExtension(filepath.Ext(filename))
			if contentType == "" {
				contentType = "application/octet-stream"
			} else {
				mediatype, _, err := mime.ParseMediaType(contentType)
				if err == nil {
					contentType = mediatype
				}
			}

			h := make(textproto.MIMEHeader)
			h.Set("Content-Disposition", `form-data; name="upfile"; filename="`+filename+`"`)
			h.Set("Content-Type", contentType)

			part, err := writer.CreatePart(h)
			if err != nil {
				t.Fatalf("Failed to create form file field: %v", err)
			}

			if _, err = io.Copy(part, file); err != nil {
				t.Fatalf("Failed to copy file data to multipart form: %v", err)
			}

			if err := writer.Close(); err != nil {
				t.Fatalf("Failed to close multipart writer: %v", err)
			}

			req, err := http.NewRequest("POST", "/api/fileanalyse", body)
			if err != nil {
				t.Fatalf("Could not created request: %v", err)
			}
			req.Header.Set("Content-Type", writer.FormDataContentType())

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(FileMetadataHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("Handler returned wrong status code: got %v want %v",
					rr.Code, http.StatusOK)
			}

			var got FileMetadata
			err = json.NewDecoder(rr.Body).Decode(&got)
			if err != nil {
				t.Fatalf("Could not decode json response: %v", err)
			}

			if got != expected {
				t.Errorf("Handler body no contain: %v got %v", expected, got)
			}
		})
	}
}
