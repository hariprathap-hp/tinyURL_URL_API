package urldomain

import "time"

type Url struct {
	OriginalURL    string    `json:"orig_url"`
	TinyURL        string    `json:"tiny_url"`
	UserEmail      string    `json:"email"`
	CreationDate   time.Time `json:"created_date"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type ListURLs struct {
	OriginalURL string `json:"orig_url"`
	TinyURL     string `json:"tiny_url"`
}

type UrlsList []ListURLs
