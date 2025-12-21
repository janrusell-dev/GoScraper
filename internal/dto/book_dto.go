package dto

type BookScrapeRequest struct {
	Page     int    `json:"page"`
	Category string `json:"category"`
	// Topic string `json:"topic,omitempty"`
}
