package scraper

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/janrusell-dev/goscraper/internal/models"
)

func BookScraper(category string, page int) ([]models.Book, error) {
	var books []models.Book
	var mu sync.Mutex

	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
		colly.CacheDir("./book_cache"),
		colly.CacheExpiration(24*time.Hour),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 4,
		Delay:       1 * time.Second,
	})

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

		mu.Lock()
		books = append(books, book)
		mu.Unlock()
	})

	slug := GetSlug(category)
	var targetURL string

	if slug == "" {
		targetURL = fmt.Sprintf("https://books.toscrape.com/catalogue/page-%d.html", page)
	} else {
		pageSegment := "index.html"

		if page > 1 {
			pageSegment = fmt.Sprintf("page-%d.html", page)
		}
		targetURL = fmt.Sprintf("https://books.toscrape.com/catalogue/category/books/%s/%s", slug, pageSegment)
	}
	log.Printf("Visiting url: %s", targetURL)
	if err := c.Visit(targetURL); err != nil {
		return nil, err
	}

	c.Wait()
	return books, nil
}
