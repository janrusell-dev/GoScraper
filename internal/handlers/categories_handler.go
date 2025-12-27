package handlers

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/janrusell-dev/goscraper/internal/scraper"
)

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories := make([]string, 0, len(scraper.CategoryMap))
	for name := range scraper.CategoryMap {
		categories = append(categories, name)
	}
	sort.Strings(categories)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
