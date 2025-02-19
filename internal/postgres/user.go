package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/http"
	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

func (c *Client) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	log := zerolog.Ctx(ctx)
	var lastInsetedId uuid.UUID

	log.Debug().Interface("user", user)

	encyptedPassword, err := http.HashPassword(user.Password)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to encrypt Password")
	}

	err = c.pool.QueryRow(ctx, "INSERT INTO users (name, email, password, phone_number, user_role) VALUES ($1,$2,$3,$4,$5) RETURNING ID", user.Name, user.Email, encyptedPassword, user.PhoneNumber, user.UserRole).Scan(&lastInsetedId)

	if err != nil {
		return nil, err
	}

	user.ID = lastInsetedId.String()

	return &user, nil
}

func (c *Client) GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	rows, err := c.pool.Query(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User
		var id uuid.UUID
		if err := rows.Scan(&id, &user.Name, &user.Email, &user.Password, &user.PhoneNumber, &user.UserRole, &user.CreatedAt); err != nil {
			return nil, err
		}
		user.ID = id.String()
		users = append(users, user)
	}

	return users, nil

}

func (c *Client) GetUserById(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	var uid uuid.UUID
	err := c.pool.QueryRow(ctx, "SELECT id, name, email, password, phone_number, user_role, created_at FROM users WHERE id = $1", id).Scan(&uid, &user.Name, &user.Email, &user.Password, &user.PhoneNumber, &user.UserRole, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	user.ID = uid.String()
	return &user, nil

}

func (c *Client) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {

	encyptedPassword, err := http.HashPassword(user.Password)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to encrypt Password")
	}
	_, err = c.pool.Exec(ctx, "UPDATE users SET name = $1, email = $2, password = $3, phone_number = $4, user_role = $5 WHERE id = $6", user.Name, user.Email, encyptedPassword, user.PhoneNumber, user.UserRole, user.ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) DeleteUser(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	var uid uuid.UUID
	err := c.pool.QueryRow(ctx, "SELECT id, name, email, password, phone_number, user_role, created_at FROM users WHERE id = $1", id).Scan(&uid, &user.Name, &user.Email, &user.Password, &user.PhoneNumber, &user.UserRole, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	user.ID = uid.String()

	_, err = c.pool.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
