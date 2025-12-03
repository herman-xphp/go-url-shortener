package services

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/xphp/go-url-shortener/internal/core/domain"
	"github.com/xphp/go-url-shortener/internal/core/ports"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type URLService struct {
	repo            ports.URLRepository
	cache           ports.CacheRepository
	shortCodeLength int
	baseURL         string
}

func NewURLService(repo ports.URLRepository, cache ports.CacheRepository, shortCodeLength int, baseURL string) *URLService {
	return &URLService{
		repo:            repo,
		cache:           cache,
		shortCodeLength: shortCodeLength,
		baseURL:         baseURL,
	}
}

func (s *URLService) ShortenURL(ctx context.Context, req *domain.CreateURLRequest) (*domain.URLResponse, error) {
	var shortCode string

	if req.CustomAlias != "" {
		existing, _ := s.repo.GetByShortCode(ctx, req.CustomAlias)
		if existing != nil {
			return nil, fmt.Errorf("custom alias already exists")
		}
		shortCode = req.CustomAlias
	} else {
		shortCode = s.generateShortCode(req.OriginalURL)
	}

	url := &domain.URL{
		OriginalURL: req.OriginalURL,
		ShortCode:   shortCode,
		CustomAlias: req.CustomAlias,
		ExpiresAt:   req.ExpiresAt,
		Clicks:      0,
	}

	if err := s.repo.Create(ctx, url); err != nil {
		return nil, fmt.Errorf("failed to create URL: %w", err)
	}

	if err := s.cache.Set(ctx, shortCode, req.OriginalURL, 3600); err != nil {
		// Log error but don't fail the request
		fmt.Printf("Failed to cache URL: %v\n", err)
	}

	return &domain.URLResponse{
		ShortURL:    fmt.Sprintf("%s/%s", s.baseURL, shortCode),
		OriginalURL: url.OriginalURL,
		ShortCode:   url.ShortCode,
		CreatedAt:   url.CreatedAt,
		ExpiresAt:   url.ExpiresAt,
	}, nil
}

func (s *URLService) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	cachedURL, err := s.cache.Get(ctx, shortCode)
	if err == nil && cachedURL != "" {
		go s.repo.IncrementClicks(context.Background(), shortCode)
		return cachedURL, nil
	}

	url, err := s.repo.GetByShortCode(ctx, shortCode)
	if err != nil {
		return "", fmt.Errorf("URL not found")
	}

	if url.ExpiresAt != nil && url.ExpiresAt.Before(url.CreatedAt) {
		return "", fmt.Errorf("URL has expired")
	}

	go s.repo.IncrementClicks(context.Background(), shortCode)

	if err := s.cache.Set(ctx, shortCode, url.OriginalURL, 3600); err != nil {
		fmt.Printf("Failed to cache URL: %v\n", err)
	}

	return url.OriginalURL, nil
}

func (s *URLService) generateShortCode(originalURL string) string {
	hash := md5.Sum([]byte(originalURL))
	hashStr := hex.EncodeToString(hash[:])
	
	num := new(big.Int)
	num.SetString(hashStr[:16], 16)
	
	return s.toBase62(num)[:s.shortCodeLength]
}

func (s *URLService) toBase62(num *big.Int) string {
	if num.Cmp(big.NewInt(0)) == 0 {
		return string(base62Chars[0])
	}

	result := ""
	base := big.NewInt(62)
	zero := big.NewInt(0)
	mod := new(big.Int)

	for num.Cmp(zero) > 0 {
		num.DivMod(num, base, mod)
		result = string(base62Chars[mod.Int64()]) + result
	}

	return result
}
