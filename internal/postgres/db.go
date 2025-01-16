package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	connStr := "user=postgres dbname=postgres password=postgres123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgressStore{
		db: db,
	}, nil
}

func (s *PostgressStore) Init() {
	s.CreateUserTable()
}

func (s *PostgressStore) CreateUserTable() string {
	return "true"
}

func (s *PostgressStore) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

func (s *PostgressStore) GetUsers(ctx context.Context) (*models.UserList, error) {
	return nil, nil

}

func (s *PostgressStore) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return nil, nil

}
func (s *PostgressStore) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil

}
func (s *PostgressStore) DeleteUser(ctx context.Context, id int) (int, error) {
	return 0, nil

}
