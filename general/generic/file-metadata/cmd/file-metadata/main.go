package main

import (
	"encoding/json"
	"github.com/dalamilla/programming-golang/general/generic/file-metadata"
	"html/template"
	"log"
	"net/http"
)

type FileMetadata struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int    `json:"size"`
}

func FormHandler(w http.ResponseWriter, req *http.Request) {
	data := map[string]string{"Title": "Form File"}
	tmpl, err := template.ParseFS(filemetadata.TemplatesFS, "templates/index.gohtml")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, data)
}

func FileMetadataHandler(w http.ResponseWriter, req *http.Request) {
	file, handler, err := req.FormFile("upfile")

	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/json")
	response := FileMetadata{Name: handler.Filename, Type: handler.Header.Get("Content-Type"), Size: int(handler.Size)}
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", FormHandler)
	mux.HandleFunc("POST /api/fileanalyse", FileMetadataHandler)

	log.Println("Server started at port 8080")
	http.ListenAndServe(":8080", mux)
}
