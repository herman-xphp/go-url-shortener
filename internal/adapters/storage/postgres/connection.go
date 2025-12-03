package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}

func CreateSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS urls (
		id BIGSERIAL PRIMARY KEY,
		original_url TEXT NOT NULL,
		short_code VARCHAR(50) UNIQUE NOT NULL,
		custom_alias VARCHAR(50) UNIQUE,
		clicks BIGINT DEFAULT 0,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		expires_at TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_short_code ON urls(short_code);
	CREATE INDEX IF NOT EXISTS idx_custom_alias ON urls(custom_alias);
	CREATE INDEX IF NOT EXISTS idx_created_at ON urls(created_at);
	`

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	return nil
}
