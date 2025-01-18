package postgres

import (
	"context"
	"database/sql"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

type Client struct {
	db *sql.DB
}

func New() (*Client, error) {
	connStr := "user=postgres dbname=postgres password=postgres123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
    return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Client{
		db: db,
	}, nil
}

func (c *Client) Init() {
	c.CreateUserTable()
}

func (c *Client) CreateUserTable() string {
	return "true"
}

func (c *Client) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

func (c *Client) GetUsers(ctx context.Context) (*models.UserList, error) {
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
