package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func loadConfigFromURL() (*pgxpool.Config, error) {
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return nil, fmt.Errorf("Must set DATABASE_URL env var")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return config, nil
}

func dbURL() (string, error) {
  dbURL, ok := os.LookupEnv("DATABASE_URL")
  if !ok {
    return "", fmt.Errorf("Must set DATABASE_URL env var")
  }

  return dbURL, nil
}

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	config, err := loadConfigFromURL()
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	return conn, nil
}

