package main

import (
  "context"

  "github.com/go-kit/kit/log"
  "gorm.io/gorm"
)

type repo struct {
  db *gorm.DB
  logger log.Logger
}

func NewRepo(db *gorm.DB, logger log.Logger) Repository {
  return &repo{
    db: db,
    logger: log.With(logger, "repo", "sql"),
  }
}

func (repo *repo) CreateUser(ctx context.Context, user User) (string, error) {
  if user.Email == "" || user.Password == "" {
    return "", nil
  }
  _ = repo.db.Create(&user)

  return "", nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
  var email string
  _ = repo.db.Raw("SELECT email FROM users WHERE id=id", id).Scan(&email)
  return email, nil
}

func (repo *repo) UpdateUser(ctx context.Context, user User) (string, error) {
  _ = repo.db.Raw("UPDATE users as u SET c.Email=?, c.Password=? WHERE u.Id = ?", user.Email, user.Password, user.Id)
  return "successfully updated", nil
}

func(repo *repo) DeleteUser(ctx context.Context, id string) (string, error) {
  _ = repo.db.Exec("DELETE FROM users WHERE id= ?", id)
  return "Succesfully deleted", nil
}
