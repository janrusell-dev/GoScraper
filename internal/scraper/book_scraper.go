package scraper

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/janrusell-dev/goscraper/internal/models"
)

func BookScraper(page int) ([]models.Book, error) {
	var books []models.Book

	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
		colly.CacheDir("./book_cache"),
		colly.CacheExpiration(24*time.Hour),
		colly.Async(true),
	)

	c.OnHTML(".product_pod", func(h *colly.HTMLElement) {
		ratingClass := h.ChildAttr(".star-rating", "class")
		rating := strings.ReplaceAll(ratingClass, "star-rating", "")
		rating = strings.TrimSpace(rating)
		availability := strings.ToLower(strings.TrimSpace(h.ChildText(".instock.availability")))
		inStock := strings.Contains(availability, "in stock")

		book := models.Book{
			Title:    h.ChildText("h3 a"),
			Price:    h.ChildText(".price_color"),
			Rating:   rating,
			InStock:  inStock,
			URL:      h.Request.AbsoluteURL(h.ChildAttr("h3 a", "href")),
			ImageURL: h.Request.AbsoluteURL(h.ChildAttr("img", "src")),
		}
		books = append(books, book)
	})

	url := fmt.Sprintf("https://books.toscrape.com/catalogue/page-%d.html", page)
	log.Printf("Visiting url: %s", url)
	if err := c.Visit(url); err != nil {
		return nil, err
	}

	c.Wait()
	return books, nil
}
