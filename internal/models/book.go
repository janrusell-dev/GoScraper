package models

type Book struct {
	Title    string `json:"title"`
	Price    string `json:"price"`
	Rating   string `json:"ratings"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	InStock  bool   `json:"in_stock"`
}
