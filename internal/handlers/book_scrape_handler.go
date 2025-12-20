package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/janrusell-dev/goscraper/internal/dto"
)

func BookScrapeHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.BookScrapeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

}
