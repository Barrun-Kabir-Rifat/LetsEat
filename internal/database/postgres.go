package database

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	dbURL := &url.URL{
		Scheme: "postgres",
		Host: net.JoinHostPort(
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
		),
		User: url.UserPassword(
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
		),
		Path: "/" + os.Getenv("DB_NAME"),
	}

	query := dbURL.Query()
	query.Set("sslmode", os.Getenv("DB_SSLMODE"))
	dbURL.RawQuery = query.Encode()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL.String())
	if err != nil {
		return nil, fmt.Errorf(
			"create database pool: %w", err,
		)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf(
			"connect to database: %w", err,
		)
	}

	return pool, nil
}
