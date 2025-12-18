package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// scraper.CourseraScraper()
	router := http.NewServeMux()
	router.HandleFunc("/hello", returnsHello)

	server := http.Server{
		Addr:    ":4040",
		Handler: router,
	}
	fmt.Println("Server running at http://localhost:4040")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func returnsHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("asdasdas"))
}
