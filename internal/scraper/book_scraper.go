package scraper

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func BookScraper() {
	fname := "json/book_scrape.json"
	file, err := os.Create(fname)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fname, err)
		return
	}
	defer file.Close()

	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com/", "https://books.toscrape.com/"),
		colly.CacheDir("./book_cache"),
		colly.CacheExpiration(24*time.Hour),
	)
	detailCollector := c.Clone()

	courses := make([]Course, 0, 200)

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		if h.Attr("class") == "Button_1qxkboh-o_O-primary_cv02ee-o_O-md_28awn8-o_O-primaryLink_109aggg" {
			return
		}

		link := h.Attr("href")

		h.Request.Visit(link)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf("visiting %s", r.URL.String())
	})

	c.OnHTML(`a.cds-CommonCard-titleLink`, func(h *colly.HTMLElement) {
		courseURL := h.Request.AbsoluteURL(h.Attr("href"))
		log.Println("Found course link:", courseURL)
	})

	detailCollector.OnHTML(`div[id=rendered-content]`, func(h *colly.HTMLElement) {
		log.Println("Course found", h.Request.URL)
		title := h.ChildText("h1.cds-119")
		if title == "" {
			log.Println("No title found", h.Request.URL)
		}
		course := Course{
			Title:       title,
			URL:         h.Request.URL.String(),
			Description: h.ChildText("div.content"),
			Creator:     h.ChildText("a.css-126fm8q span.css-6ecy9b"),
			Rating:      h.ChildText("div.css-h1jogs"),
			Level:       h.ChildText("div.css-fk6qfz"),
		}

		h.ForEach(".AboutCourse .ProductGlance > div", func(_ int, el *colly.HTMLElement) {
			svgTitle := strings.Split(el.ChildText("div:nth-child(1) svg title"), " ")
			lastWord := svgTitle[len(svgTitle)-1]
			switch lastWord {
			case "languages":
				course.Language = el.ChildText("div:nth-child(2) > div:nth-child(1)")
			case "Level":
				course.Level = el.ChildText("div:nth-child(2) > div:nth-child(1)")
			case "complete":
				course.Commitment = el.ChildText("div:nth-child(2) > div:nth-child(1)")
			}
		})
		courses = append(courses, course)
	})

	c.Visit("https://coursera.org/browse/computer-science")

	enc := json.NewEncoder(file)
	enc.SetIndent("", " ")

	enc.Encode(courses)

}
