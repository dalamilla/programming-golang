package main

import (
	"encoding/json"
	"github.com/dalamilla/programming-golang/general/generic/system-metrics/internal/metrics"
	"log"
	"net/http"
)

func SystemMetricsHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := metrics.GetSystemMetrics()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		error := metrics.SystemMetricsError{Error: "Reading System Metrics"}
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/system", SystemMetricsHandler)

	log.Println("Server started at port 8080")
	http.ListenAndServe(":8080", mux)
}
