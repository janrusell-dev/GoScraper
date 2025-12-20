package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/janrusell-dev/goscraper/internal/dto"
)

func AvailableHandler(w http.ResponseWriter, r *http.Request) {
	availables := []dto.AvailableResponse{
		{Type: "Course", Link: "https://www.coursera.org/browse"},
		{Type: "Book", Link: "https://books.toscrape.com/"},
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(availables); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}
