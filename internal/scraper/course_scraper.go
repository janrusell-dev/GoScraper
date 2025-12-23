package scraper

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/janrusell-dev/goscraper/internal/models"
)

func CourseraScraper() {

	c := colly.NewCollector(
		colly.AllowedDomains("coursera.org", "www.coursera.org"),
		colly.CacheDir("./coursera_cache"),
		colly.CacheExpiration(24*time.Hour),
	)
	detailCollector := c.Clone()

	targetKeywords := []string{
		"backend", "python", "nodejs", "java",
		"api", "rest", "database", "sql", "mongodb",
		"django", "flask", "spring boot", "express",
		"ai", "devops", "cybersecurity", "data structures and algorithms",
		"back-end", "generative ai", "golang", "fastapi", "rust",
	}

	courses := make([]models.Course, 0, 200)

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		if h.Attr("class") == "Button_1qxkboh-o_O-primary_cv02ee-o_O-md_28awn8-o_O-primaryLink_109aggg" {
			return
		}

		link := h.Attr("href")

		if !strings.HasPrefix(link, "/browse") || strings.Index(link, "=signup") > -1 || strings.Index(link, "=login") > -1 {
			return
		}

		h.Request.Visit(link)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf("visiting %s", r.URL.String())
	})

	c.OnHTML(`a.cds-CommonCard-titleLink`, func(h *colly.HTMLElement) {
		courseURL := h.Request.AbsoluteURL(h.Attr("href"))
		title := h.ChildText("h3.cds-CommonCard-title")
		log.Println("Found course link:", courseURL)
		if containsAny(strings.ToLower(title), targetKeywords) {
			log.Printf("Found course: %s", title)
			if strings.Contains(courseURL, "coursera.org/learn") {
				detailCollector.Visit(courseURL)
			}
		}
	})

	detailCollector.OnHTML(`div[id=rendered-content]`, func(h *colly.HTMLElement) {
		log.Println("Course found", h.Request.URL)
		title := h.ChildText("h1.cds-119")
		if title == "" {
			log.Println("No title found", h.Request.URL)
		}
		course := models.Course{
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

}

func containsAny(s string, targetKeywords []string) bool {
	for _, keyword := range targetKeywords {
		if strings.Contains(s, keyword) {
			return true
		}
	}
	return false
}
