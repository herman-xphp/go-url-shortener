package ports

import (
	"context"

	"github.com/xphp/go-url-shortener/internal/core/domain"
)

type URLRepository interface {
	Create(ctx context.Context, url *domain.URL) error
	GetByShortCode(ctx context.Context, shortCode string) (*domain.URL, error)
	GetByID(ctx context.Context, id int64) (*domain.URL, error)
	IncrementClicks(ctx context.Context, shortCode string) error
	Delete(ctx context.Context, shortCode string) error
}

type CacheRepository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, ttl int) error
	Delete(ctx context.Context, key string) error
}
