package repositories

import (
	"context"

	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/models"
)

type UserRepository interface {
	// define crud functions
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	GetUserById(ctx context.Context, id int) (*models.User, error)
	// TODO: GetFilteredUser()
	UpdateUser(ctx context.Context, id int, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int) (int, error)
}

type TestUserRepo struct {
	Users []models.User
}

func (repo *TestUserRepo) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	repo.Users = append(repo.Users, user)
	return repo.Users[len(repo.Users)-1], nil
}

func (repo *TestUserRepo) GetUsers(ctx context.Context) []models.User {
	return repo.Users
}

func (repo *TestUserRepo) GetUserById(ctx context.Context, id int) (models.User, error) {
	return repo.Users[id], nil
}

func (repo *TestUserRepo) UpdateUser(ctx context.Context, id int, user models.User) (models.User, error) {
	repo.Users[id] = user
	return repo.Users[id], nil
}

func (repo *TestUserRepo) DeleteUser(ctx context.Context, id int) (int, error) {
	repo.Users = append(repo.Users[:id], repo.Users[id+1:]...)
	return id, nil
}
