package postgres

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

func (c *Client) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

func (c *Client) GetUsers(ctx context.Context) ([]models.User, error) {
	log := zerolog.Ctx(ctx)

	rows, err := c.pool.Exec(ctx, "SELECT * FROM users")
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to get users from the DB")
	}
	log.Print(rows)
	return nil, nil

}

func (c *Client) GetUserById(ctx context.Context, id string) (*models.User, error) {
	return nil, nil

}
func (c *Client) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil

}
func (c *Client) DeleteUser(ctx context.Context, id string) (int, error) {
	return 0, nil

}
