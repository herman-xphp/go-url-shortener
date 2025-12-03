package domain

import "time"

type URL struct {
	ID           int64     `json:"id" db:"id"`
	OriginalURL  string    `json:"original_url" db:"original_url"`
	ShortCode    string    `json:"short_code" db:"short_code"`
	CustomAlias  string    `json:"custom_alias,omitempty" db:"custom_alias"`
	Clicks       int64     `json:"clicks" db:"clicks"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty" db:"expires_at"`
}

type CreateURLRequest struct {
	OriginalURL string     `json:"original_url" validate:"required,url"`
	CustomAlias string     `json:"custom_alias,omitempty" validate:"omitempty,alphanum,min=4,max=20"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
}

type URLResponse struct {
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
}

type URLStats struct {
	URL         *URL  `json:"url"`
	TotalClicks int64 `json:"total_clicks"`
}
