package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

func (c *Client) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	log := zerolog.Ctx(ctx)
	var lastInsetedId uuid.UUID
	log.Print(user)
	err := c.pool.QueryRow(ctx, "INSERT INTO users (name, email, password, phone_number, user_role) VALUES ($1,$2,$3,$4,$5) RETURNING ID", user.Name, user.Email, user.Password, user.PhoneNumber, user.UserRole).Scan(&lastInsetedId)

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to insert user to db")
		return nil, err
	}

	user.ID = lastInsetedId.String()

	return &user, nil
}

func (c *Client) GetUsers(ctx context.Context) ([]models.User, error) {
	log := zerolog.Ctx(ctx)
	var users []models.User
	rows, err := c.pool.Query(ctx, "SELECT * FROM users")
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to get users from the DB")
	}

	var id uuid.UUID

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&id, &user.Name, &user.Email, &user.Password, &user.PhoneNumber, &user.UserRole, &user.CreatedAt); err != nil {
			log.Error().
				Err(err).
				Msg("failed to scan user into users")
			return nil, err
		}
		user.ID = id.String()

		users = append(users, user)

	}

	return users, nil

}

func (c *Client) GetUserById(ctx context.Context, id string) (*models.User, error) {
	log := zerolog.Ctx(ctx)
	var user models.User
	var uuidId uuid.UUID
	err := c.pool.QueryRow(ctx, "SELECT id, name, email, password, phone_number, user_role, created_at FROM users WHERE id = $1", id).Scan(&uuidId, &user.Name, &user.Email, &user.Password, &user.PhoneNumber, &user.UserRole, &user.CreatedAt)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to get user from the DB")
		return nil, err
	}

	user.ID = uuidId.String()
	return &user, nil

}
func (c *Client) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	log := zerolog.Ctx(ctx)
	_, err := c.pool.Exec(ctx, "UPDATE users SET name = $1, email = $2, password = $3, phone_number = $4, user_role = $5 WHERE id = $6", user.Name, user.Email, user.Password, user.PhoneNumber, user.UserRole, user.ID)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to update user in the DB")
		return nil, err
	}
	return &user, nil

}
func (c *Client) DeleteUser(ctx context.Context, id string) (*models.User, error) {
	log := zerolog.Ctx(ctx)
	var user models.User
	var uuidId uuid.UUID
	err := c.pool.QueryRow(ctx, "SELECT id, name, email, password, phone_number, user_role, created_at FROM users WHERE id = $1", id).Scan(&uuidId, &user.Name, &user.Email, &user.Password, &user.PhoneNumber, &user.UserRole, &user.CreatedAt)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to get user from the DB")
		return nil, err
	}

	user.ID = uuidId.String()

	_, err = c.pool.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to delete user from the DB")
		return nil, err
	}

	return &user, nil

}
