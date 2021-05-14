package usersvc

import (
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, user User) (string, error)
	DeleteUser(ctx context.Context, id string) (string, error)
}
