package usersvc

import (
  "context"
)

type User struct {
  Id string `gorm:"primaryKey" json:"id, omitempty"`
  Email string `json:"email" gorm:"type:text"`
  Password string `json:"password"`
}

type Repository interface {
  CreateUser(ctx context.Context, user User) (string, error)
  GetUser(ctx context.Context, id string) (string, error)
  GetAllUsers(ctx context.Context) ([]User, error)
  UpdateUser(ctx context.Context, user User) (string, error)
  DeleteUser(ctx context.Context, id string) (string, error)
}
