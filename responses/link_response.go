package responses

import "time"

type LinkResponse struct {
	Code      string     `json:"code"`
	ExpiresAt *time.Time `json:"expiresAt"`
}

type UrlResponse struct {
	Url string `json:"url"`
}
