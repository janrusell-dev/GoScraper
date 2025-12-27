package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/janrusell-dev/goscraper/internal/dto"
	"github.com/janrusell-dev/goscraper/internal/scraper"
)

func BookScrapeHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.BookScrapeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	books, err := scraper.BookScraper(req.Category, req.Page)
	if err != nil {
		http.Error(w, "Failed to scrape books", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}
