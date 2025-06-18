package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Client struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context) (*Client, error) {
	log := zerolog.Ctx(ctx)

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return nil, fmt.Errorf("DATABASE_URL not set")
	}
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	log.Info().Msg("Connected to postgres database")

	return &Client{
		pool: pool,
	}, nil
}

func (c *Client) Shutdown() {
	c.pool.Close()
}
