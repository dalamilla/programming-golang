package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Timestamp struct {
	Unix int64  `json:"unix"`
	UTC  string `json:"utc"`
}

type TimestampError struct {
	Error string `json:"error"`
}

func parsingDate(strDate string) (time.Time, error) {
	validFormats := []string{
		"2006-01-02",           // %Y-%m-%d
		"02 January 2006",      // %d %B %Y
		"02 January 2006, MST", // %d %B %Y, %Z
	}

	for _, fmt := range validFormats {
		t, err := time.Parse(fmt, strDate)
		if err == nil {
			return t.UTC(), nil
		}
	}
	return time.Time{}, strconv.ErrSyntax
}

func TimestampHandler(w http.ResponseWriter, req *http.Request) {
	inputDate := req.PathValue("input_date")
	var timeDate time.Time

	if inputDate == "" {
		timeDate = time.Now().UTC()
	} else {
		ms, err := strconv.ParseInt(inputDate, 10, 64)
		if err == nil {
			timeDate = time.UnixMilli(ms).UTC()
		} else {
			timeDate, err = parsingDate(inputDate)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				error := TimestampError{Error: "Invalid Date"}
				json.NewEncoder(w).Encode(error)
				return
			}
		}
	}

	unix := timeDate.UnixMilli()
	utc := timeDate.In(time.FixedZone("GMT", 0)).Format(time.RFC1123)

	w.Header().Set("Content-Type", "application/json")
	response := Timestamp{Unix: unix, UTC: utc}
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := AppMux()
	log.Println("Server started at port 8080")
	http.ListenAndServe(":8080", mux)
}

func AppMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api", TimestampHandler)
	mux.HandleFunc("GET /api/", TimestampHandler)
	mux.HandleFunc("GET /api/{input_date}", TimestampHandler)
	return mux
}
