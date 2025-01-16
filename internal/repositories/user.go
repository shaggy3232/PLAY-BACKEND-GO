package repositories

import (
	"context"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

type UserRepository interface {
	// define crud functions
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	GetUsers(ctx context.Context) (*models.UserList, error)
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	// TODO: GetFilteredUser()
	UpdateUser(ctx context.Context, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int) (int, error)
}
