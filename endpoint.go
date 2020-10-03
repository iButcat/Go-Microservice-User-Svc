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
    msg, err := service.UpdateUser(ctx, req.user)
    return UpdateUserResponse{
      Msg: msg,
      Err: err,
    }, err
  }
}

func MakeDeleteUserEndpoint(service Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    req := request.(DeleteUserRequest)
    msg, err := service.DeleteUser(ctx, req.Id)
    if err != nil {
      return DeleteUserResponse{
        Msg: msg,
        Err: err,
      }, nil
    }
    return DeleteUserResponse{
      Msg: msg,
      Err: err,
    }, nil
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
    user User
  }

  UpdateUserResponse struct {
    Msg string `json:"response"`
    Err error `json:"error, omitempty"`
  }

  DeleteUserRequest struct {
    Id string `json:"id"`
  }

  DeleteUserResponse struct {
    Msg string `json:"response"`
    Err error `json:"error, omitempty"`
  }
)
