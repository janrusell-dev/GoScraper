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
	router.HandleFunc("GET /", handlers.IndexHandler)

	scrapeStack := middleware.CreateStack(
		middleware.RateLimiter,
	)

	router.Handle("POST /scrape/books", scrapeStack(http.HandlerFunc(handlers.BookScrapeHandler)))
	router.Handle("GET /scraped/books", scrapeStack(http.HandlerFunc(handlers.ScrapedBooksHandler)))
	router.Handle("POST /scrape/course", scrapeStack(http.HandlerFunc(handlers.CourseHandler)))
	router.Handle("GET /scraped/course", scrapeStack(http.HandlerFunc(handlers.ScrapedCourseHandler)))
	router.Handle("GET /available", scrapeStack(http.HandlerFunc(handlers.AvailableHandler)))
	router.Handle("GET /books/categories", scrapeStack(http.HandlerFunc(handlers.CategoriesHandler)))

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))
	stack := middleware.CreateStack(
		middleware.CorsMiddleware,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(v1),
	}
	fmt.Println("Server running at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
