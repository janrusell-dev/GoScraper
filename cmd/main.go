package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/janrusell-dev/goscraper/internal/handlers"
	"github.com/janrusell-dev/goscraper/internal/middleware"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /health", handlers.HealthCheckHandler)
	router.HandleFunc("POST /scrape/books", handlers.BookScrapeHandler)
	router.HandleFunc("GET /scraped/books", handlers.ScrapedBooksHandler)
	router.HandleFunc("POST /scrape/course", handlers.CourseHandler)
	router.HandleFunc("GET /scraped/course", handlers.ScrapedCourseHandler)
	router.HandleFunc("GET /available", handlers.AvailableHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.CorsMiddleware(router),
	}
	fmt.Println("Server running at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
