package postgres

import (
	"database/sql"
	"fmt"

	// "github.com/google/uuid"

	_ "github.com/lib/pq"
	// "bot/storage"
)

type Storage struct {
	db *sql.DB
}

// New creates new PostgreSQL storage.
func New(connStr string) (*Storage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	return &Storage{db: db}, nil
}