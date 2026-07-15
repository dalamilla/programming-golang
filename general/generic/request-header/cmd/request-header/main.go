package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

type WhoAmI struct {
	Software  string `json:"software"`
	Language  string `json:"language"`
	IpAddress string `json:"ipaddress"`
}

func WhoAmIHandler(w http.ResponseWriter, req *http.Request) {
	userAgent := req.Header.Get("User-Agent")
	acceptLanguage := req.Header.Get("Accept-Language")
	ipAddress, _, _ := net.SplitHostPort(req.RemoteAddr)

	w.Header().Set("Content-Type", "application/json")
	response := WhoAmI{Software: userAgent, Language: acceptLanguage, IpAddress: ipAddress}
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/whoami", WhoAmIHandler)

	log.Println("Server started at port 8080")
	http.ListenAndServe(":8080", mux)
}
