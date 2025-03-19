package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

func (c *Client) CreateBooking(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	var bookingID uuid.UUID

	err := c.pool.QueryRow(ctx, "INSERT INTO bookings (referee_id, organizer_id, price, start_time, end_time, location, accepted, cancelled) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING ID", booking.RefereeID, booking.OrganizerID, booking.Price, booking.Start, booking.End, booking.Location, false, booking.Cancelled).Scan(&bookingID)

	if err != nil {
		return nil, err
	}

	booking.ID = bookingID.String()

	return booking, nil
}

func (c *Client) GetBookings(ctx context.Context) ([]models.Booking, error) {
	var bookings []models.Booking

	rows, err := c.pool.Query(ctx, "SELECT * FROM bookings")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var booking models.Booking
		var id uuid.UUID

		if err := rows.Scan(&id, &booking.RefereeID, &booking.OrganizerID, &booking.Price, &booking.Start, &booking.End, &booking.Location, &booking.Accepted, &booking.Cancelled, &booking.LastUpdated, &booking.CreatedAt); err != nil {
			return nil, err
		}
		booking.ID = id.String()
		bookings = append(bookings, booking)

	}

	return bookings, nil
}

func (c *Client) GetBookingById(ctx context.Context, id string) (*models.Booking, error) {
	var booking models.Booking
	var bookingID uuid.UUID

	err := c.pool.QueryRow(ctx, "SELECT id, referee_id, organizer_id, price, start_time, end_time, location, accepted, cancelled, last_updated, created_at FROM bookings WHERE id = $1", id).Scan(&bookingID, &booking.RefereeID, &booking.OrganizerID, &booking.Price, &booking.Start, &booking.End, &booking.Location, &booking.Accepted, &booking.Cancelled, &booking.LastUpdated, &booking.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &booking, err
}

func (c *Client) DeleteBooking(ctx context.Context, id string) (*models.Booking, error) {
	booking, err := c.GetBookingById(ctx, id)

	if err != nil {
		return nil, err
	}

	_, err = c.pool.Exec(ctx, "DELETE FROM bookings WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (c *Client) CheckConflicts(ctx context.Context, userID string, start time.Time, end time.Time) (bool, error) {

	return false, nil
}

func (c *Client) AcceptBooking(ctx context.Context, id string) (*models.Booking, error) {

	return nil, nil
}

func (c *Client) EditBooking(ctx context.Context, booking models.Booking) (models.Booking, error) {
	var updatedBooking models.Booking

	err := c.pool.QueryRow(ctx, "UPDATE bookings SET referee_id = $1, organizer_id = $2, price = $3, start_time = $4, end_time = $5, location = $6, accepted = $7, cancelled = $8, last_updated = $9 WHERE id = $10 RETURNING *", booking.RefereeID, booking.OrganizerID, booking.Price, booking.Start, booking.End, booking.Location, booking.Accepted, booking.Cancelled, booking.CreatedAt, time.Now(), booking.ID).Scan(updatedBooking)
	if err != nil {
		return updatedBooking, err
	}

	return updatedBooking, nil
}
