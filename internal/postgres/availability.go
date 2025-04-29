package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

func (c *Client) CreateAvailability(ctx context.Context, newAvailability *models.Availability) (*models.Availability, error) {
	log := zerolog.Ctx(ctx)
	var availID uuid.UUID

	err := c.pool.QueryRow(ctx, "INSERT INTO availabilities (user_id, price, start_time, end_time) VALUES ($1,$2,$3,$4) RETURNING id", newAvailability.UserID, newAvailability.Price, newAvailability.Start, newAvailability.End).Scan(&availID)
	if err != nil {
		log.Error().
			Err(err).
			Msg("cannot insert into database")
		return nil, err
	}

	newAvailability.ID = availID.String()

	return newAvailability, nil
}

func (c *Client) GetAvailabilities(ctx context.Context) ([]models.Availability, error) {
	var availabilities []models.Availability

	rows, err := c.pool.Query(ctx, "SELECT * FROM availabilities")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var avail models.Availability
		var id uuid.UUID
		var user uuid.UUID

		if err := rows.Scan(&id, &user, &avail.Price, &avail.Start, &avail.End, &avail.CreatedAt); err != nil {
			fmt.Println(err)
			return nil, err
		}
		avail.ID = id.String()
		avail.UserID = user.String()

		availabilities = append(availabilities, avail)
	}

	return availabilities, nil
}

func (c *Client) GetAvailabilityById(ctx context.Context, id string) (*models.Availability, error) {
	var avail models.Availability
	var availID uuid.UUID
	var userID uuid.UUID

	err := c.pool.QueryRow(ctx, "SELECT id, user_id, price, start_time, end_time from availabilities WHERE id = $1", id).Scan(&availID, &userID, &avail.Price, &avail.Start, &avail.End)

	if err != nil {
		return nil, err
	}

	avail.ID = availID.String()
	avail.UserID = userID.String()

	return &avail, nil
}

func (c *Client) GetAvailabilityByUser(ctx context.Context, userID string) ([]models.Availability, error) {
	log := zerolog.Ctx(ctx)
	var availabilities []models.Availability

	rows, err := c.pool.Query(ctx, "SELECT * FROM availabilities WHERE user_id = $1", userID)

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to get get rows for data base")
		return nil, err

	}

	for rows.Next() {
		var avail models.Availability
		var availID uuid.UUID
		var userID uuid.UUID

		if err := rows.Scan(&availID, &userID, &avail.Price, &avail.Start, &avail.End, &avail.CreatedAt); err != nil {
			log.Error().
				Err(err).
				Msg("failed to scan row from query")
			return nil, err
		}

		avail.ID = availID.String()
		avail.UserID = userID.String()
		availabilities = append(availabilities, avail)
	}

	return availabilities, nil
}

func (c *Client) UpdateAvailability(ctx context.Context, avaialbility models.Availability) (*models.Availability, error) {
	var updatedAvialID uuid.UUID
	var updatedAvail models.Availability

	err := c.pool.QueryRow(ctx, "UPDATE availabilities SET price = $1, start_time = $2, end_time = $3 WHERE id = $4 RETURNING id, user_id, price, start_time, end_time, created_at", avaialbility.Price, avaialbility.Start, avaialbility.End, avaialbility.ID).Scan(&updatedAvialID, &updatedAvail.UserID, &updatedAvail.Price, &updatedAvail.Start, &updatedAvail.Start, &updatedAvail.End)

	if err != nil {
		return nil, err
	}
	updatedAvail.ID = updatedAvialID.String()
	return &updatedAvail, nil

}

func (c *Client) DeleteAvailability(ctx context.Context, id string) (*models.Availability, error) {
	var avail models.Availability
	var availID uuid.UUID
	log := zerolog.Ctx(ctx)

	log.Print(id)
	err := c.pool.QueryRow(ctx, "SELECT * FROM availabilities WHERE id = $1", id).Scan(&availID, &avail.UserID, &avail.Price, &avail.Start, &avail.End, &avail.CreatedAt)
	if err != nil {
		return nil, err
	}
	avail.ID = availID.String()

	_, err = c.pool.Exec(ctx, "DELETE FROM availabilities WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &avail, nil
}

func (c *Client) GetValidAvailabilities(ctx context.Context, start time.Time, end time.Time) ([]models.Availability, error) {
	log := zerolog.Ctx(ctx)
	var availabilities []models.Availability

	rows, err := c.pool.Query(ctx, `
		SELECT * FROM availabilities
		WHERE start_time <= $1 AND end_time >= $2
	`, start, end)

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to get availabilities within time range")
		return nil, err
	}

	for rows.Next() {
		var avail models.Availability
		var availID uuid.UUID
		var userID uuid.UUID

		if err := rows.Scan(&availID, &userID, &avail.Price, &avail.Start, &avail.End, &avail.CreatedAt); err != nil {
			log.Error().
				Err(err).
				Msg("failed to scan availability row")
			return nil, err
		}

		avail.ID = availID.String()
		avail.UserID = userID.String()
		availabilities = append(availabilities, avail)
	}

	return availabilities, nil
}
