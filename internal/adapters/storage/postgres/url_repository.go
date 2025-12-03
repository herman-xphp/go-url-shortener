package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/herman-xphp/go-url-shortener/internal/core/domain"
	"github.com/herman-xphp/go-url-shortener/internal/core/ports"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) ports.URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(ctx context.Context, url *domain.URL) error {
	query := `
		INSERT INTO urls (original_url, short_code, custom_alias, clicks, created_at, expires_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	url.CreatedAt = time.Now()
	url.Clicks = 0

	err := r.db.QueryRowContext(
		ctx,
		query,
		url.OriginalURL,
		url.ShortCode,
		url.CustomAlias,
		url.Clicks,
		url.CreatedAt,
		url.ExpiresAt,
	).Scan(&url.ID, &url.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create URL: %w", err)
	}

	return nil
}

func (r *URLRepository) GetByShortCode(ctx context.Context, shortCode string) (*domain.URL, error) {
	query := `
		SELECT id, original_url, short_code, custom_alias, clicks, created_at, expires_at
		FROM urls
		WHERE short_code = $1 OR custom_alias = $1
	`

	url := &domain.URL{}
	err := r.db.QueryRowContext(ctx, query, shortCode).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortCode,
		&url.CustomAlias,
		&url.Clicks,
		&url.CreatedAt,
		&url.ExpiresAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("URL not found")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get URL: %w", err)
	}

	return url, nil
}

func (r *URLRepository) GetByID(ctx context.Context, id int64) (*domain.URL, error) {
	query := `
		SELECT id, original_url, short_code, custom_alias, clicks, created_at, expires_at
		FROM urls
		WHERE id = $1
	`

	url := &domain.URL{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortCode,
		&url.CustomAlias,
		&url.Clicks,
		&url.CreatedAt,
		&url.ExpiresAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("URL not found")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get URL: %w", err)
	}

	return url, nil
}

func (r *URLRepository) IncrementClicks(ctx context.Context, shortCode string) error {
	query := `
		UPDATE urls
		SET clicks = clicks + 1
		WHERE short_code = $1 OR custom_alias = $1
	`

	_, err := r.db.ExecContext(ctx, query, shortCode)
	if err != nil {
		return fmt.Errorf("failed to increment clicks: %w", err)
	}

	return nil
}

func (r *URLRepository) Delete(ctx context.Context, shortCode string) error {
	query := `DELETE FROM urls WHERE short_code = $1 OR custom_alias = $1`

	_, err := r.db.ExecContext(ctx, query, shortCode)
	if err != nil {
		return fmt.Errorf("failed to delete URL: %w", err)
	}

	return nil
}
