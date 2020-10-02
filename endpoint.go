package main

import (
  "context"

  "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
  CreateUserEndpoint endpoint.Endpoint
  GetUserEndpoint endpoint.Endpoint
  UpdateUserEndpoint endpoint.Endpoint
  DeleteUserEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
  return Endpoints{
    CreateUserEndpoint: MakeCreateUserEndpoint(s),
    GetUserEndpoint: MakeGetUserEndpoint(s),
    UpdateUserEndpoint: MakeUpdateUserEndpoint(s),
    DeleteUserEndpoint: MakeDeleteUserEndpoint(s),
  }
}

func MakeCreateUserEndpoint(service Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    req := request.(CreateUserRequest)
    ok, err := service.CreateUser(ctx, req.Email, req.Password)
    return CreateUserResponse{
      V: ok,
      }, err
  }
}

func MakeGetUserEndpoint(service Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    req := request.(GetUserRequest)
    email, err := service.GetUser(ctx, req.Id)
    return GetUserResponse{
      Email: email,
    }, err
  }
}

func MakeUpdateUserEndpoint(service Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    req := request.(UpdateUserRequest)
    email, err := service.UpdateUser(ctx, req.Id)
    return UpdateUserResponse{
      Email: email,
    }, err
  }
}

func MakeDeleteUserEndpoint(service Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    req := request.(DeleteUserRequest)
    email, err := service.DeleteUser(ctx, req.Id)
    return DeleteUserResponse{
      Email: email,
    }
  }
}

type (
  CreateUserRequest struct {
    Email string `json:"email"`
    Password string `json:"password"`
  }

  CreateUserResponse struct {
    V string `json:"v"`
    Err error `json:"-"`
  }

  GetUserRequest struct {
    Id string `json:"id"`
  }

  GetUserResponse struct {
    Email string `json:"email"`
    Err error `json:"-"`
  }

  UpdateUserRequest struct {
    Id string `json:"id"`
  }

  UpdateUserResponse struct {
    Email string `json:"email"`
    Err error `json:"-"`
  }

  DeleteUserRequest struct {
    Id string `json:"id"`
  }

  DeleteUserResponse struct {
    Email string `json:"email"`
  }
)
