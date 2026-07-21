package main

import (
	"github.com/dalamilla/programming-golang/general/generic/url-shortener/internal/database"
	"github.com/dalamilla/programming-golang/general/generic/url-shortener/internal/handler"
	"github.com/dalamilla/programming-golang/general/generic/url-shortener/internal/repository"
	"log"
	"net/http"
)

func main() {
	log.Println("Initialize db")
	db, err := database.InitDB("app.db")
	if err != nil {
		log.Fatalf("Failed to initialize BoltDB: %v", err)
	}
	defer db.Close()

	log.Println("Initialize repository")
	shorterRepo, err := repository.NewShorterRepository(db)
	if err != nil {
		log.Fatalf("Failed to initialize user repository: %v", err)
	}

	log.Println("Initialize handlers")
	shorterHandler := handler.NewShorterHandler(shorterRepo)

	log.Println("Setup routes")
	mux := AppMux(shorterHandler)

	log.Println("Server started at port 8080")
	http.ListenAndServe(":8080", mux)
}

func AppMux(shorterHandler *handler.ShorterHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/shorturl", shorterHandler.CreateURLShortenerHandler)
	mux.HandleFunc("GET /api/shorturl", shorterHandler.GetURLShortenerHandler)
	mux.HandleFunc("GET /api/shorturl/", shorterHandler.GetURLShortenerHandler)
	mux.HandleFunc("GET /api/shorturl/{id}", shorterHandler.GetURLShortenerHandler)
	return mux
}
