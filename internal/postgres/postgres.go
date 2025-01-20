package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Client struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, user string, password string, host string, name string) (*Client, error) {
	log := zerolog.Ctx(ctx)

	url := fmt.Sprintf("postgres://%s:%s@%s/%s", user, password, host, name)
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
