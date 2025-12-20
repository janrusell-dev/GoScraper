package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/janrusell-dev/goscraper/internal/handlers"
)

func main() {
	// scraper.CourseraScraper()
	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.AvailableHandler)
	router.HandleFunc("GET /health", handlers.HealthCheckHandler)
	router.HandleFunc("POST /scrape/books", handlers.BookScrapeHandler)
	router.HandleFunc("GET /scraped/books", handlers.ScrapedBooksHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("Server running at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
