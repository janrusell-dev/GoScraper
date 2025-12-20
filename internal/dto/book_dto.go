package dto

type BookRequestDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

type BookResponseDTO struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
	Publisher   string `json:"publisher"`
	PublishedAt string `json:"published_at"`
}
