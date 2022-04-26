package urldomain

import (
	"strings"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
	"time"
)

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

type UniqKeys struct {
	Keys []string `json:"keys"`
}

type UrlsList []ListURLs

func (url *Url) Validate() *errors.RestErr {
	url.OriginalURL = strings.TrimSpace(strings.ToLower(url.OriginalURL))
	if url.OriginalURL == "" {
		return errors.NewBadRequestError("invalid original url input. should not be empty")
	}
	url.UserEmail = strings.TrimSpace(strings.ToLower(url.UserEmail))
	if url.OriginalURL == "" {
		return errors.NewBadRequestError("invalid email input. should not be empty")
	}
	return nil
}
