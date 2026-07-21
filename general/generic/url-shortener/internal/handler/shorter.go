package handler

import (
	"encoding/json"
	"github.com/dalamilla/programming-golang/general/generic/url-shortener/internal/repository"
	"github.com/dalamilla/programming-golang/general/generic/url-shortener/internal/types"
	"net/http"
	"strconv"
	"strings"
)

type ShorterHandler struct {
	repo *repository.ShorterRepository
}

func NewShorterHandler(repo *repository.ShorterRepository) *ShorterHandler {
	return &ShorterHandler{repo: repo}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(types.ShorterError{Error: message})
}

func (h *ShorterHandler) CreateURLShortenerHandler(w http.ResponseWriter, req *http.Request) {
	var payload types.ShorterPayload

	contentType := strings.Split(req.Header.Get("Content-Type"), ";")[0]

	switch contentType {
	case "application/json":
		json.NewDecoder(req.Body).Decode(&payload)
	case "application/x-www-form-urlencoded":
		req.ParseForm()
		payload.URL = req.FormValue("url")
	default:
		respondWithError(w, http.StatusUnsupportedMediaType, "Unsupported Content-type")
		return
	}

	err := payload.Validate()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.repo.Create(payload.URL)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Can't create url record")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ShorterHandler) GetURLShortenerHandler(w http.ResponseWriter, req *http.Request) {
	inputID := req.PathValue("id")

	if inputID == "" {
		respondWithError(w, http.StatusNotFound, "No short URL found for the given input")
		return
	}

	ID, err := strconv.ParseUint(inputID, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Wrong format")
		return
	}

	response, err := h.repo.Get(ID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "No short URL found for the given input")
		return
	}
	http.Redirect(w, req, response.OriginalURL, http.StatusFound)
}
