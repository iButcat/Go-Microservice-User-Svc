package main

import (
  "context"

  "github.com/gofrs/uuid"
  "github.com/go-kit/kit/log"
)

type Service interface {
  CreateUser(ctx context.Context, email string, password string) (string, error)
  GetUser(ctx context.Context, id string) (string, error)
  UpdateUser(ctx context.Context, id string) (string, error)
  DeleteUser(ctx context.Context, id string) (string, error)
}

type basicService struct{
  repository Repository
  logger log.Logger
}

func NewBasicService(repo Repository, logger log.Logger) Service {
  return &basicService{
    repository: repo,
    logger: logger,
  }
}

func (b basicService) CreateUser(ctx context.Context, email string, password string) (string, error) {
  logger := log.With(b.logger, "method", "CreateUser")
  uuid, _ := uuid.NewV4()
  id := uuid.String()
  user := User{
    Id: id,
    Email: email,
    Password: password,
  }
  if _, err := b.repository.CreateUser(ctx, user); err != nil {
    return "", err
  }
  logger.Log("create user", id)
  return "Success", nil
}

func (b basicService) GetUser(ctx context.Context, id string) (string, error) {
  logger := log.With(b.logger, "method", "GetUser")
  email, err := b.repository.GetUser(ctx, id)
  if err != nil {
    return "", err
  }
  logger.Log("get user", id)
  return email, nil
}

func (b basicService) UpdateUser(ctx context.Context, id string) (string, error) {
  logger := log.With(b.logger, "method", "UpdateUser")
  email, err := b.repository.UpdateUser(ctx, id)
  if err != nil {
    return "", err
  }
  logger.Log("update user", id)
  return email, nil
}

func (b basicService) DeleteUser(ctx context.Context, id string) (string, error) {
  logger := log.With(b.logger, "method", "DeleteUser")
  id, err := b.repository.DeleteUser(ctx, id)
  if err != nil {
    return "", err
  }
  logger.Log("delete user", id)
  return "", nil
}
