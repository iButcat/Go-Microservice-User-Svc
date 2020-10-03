package main

import (
  "context"
  "net/http"
  "encoding/json"

  "github.com/go-kit/kit/log"
  "github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"

)

func MakeHTTPHandler(service Service, logger log.Logger) http.Handler {
  r := mux.NewRouter()
  e := MakeServerEndpoints(service)

  r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
    e.CreateUserEndpoint,
    decodeCreateProfileRequest,
    encodeResponse,
    ))

  r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
    e.GetUserEndpoint,
    decodeGetUserRequest,
    encodeResponse,
    ))

  r.Methods("PUT").Path("/user/{id}").Handler(httptransport.NewServer(
    e.UpdateUserEndpoint,
    decodeUpdateUserRequest,
    encodeResponse,
    ))

  r.Methods("DELETE").Path("/user/{id}").Handler(httptransport.NewServer(
    e.DeleteUserEndpoint,
    decodeDeleteUserRequest,
    encodeResponse,
    ))

    return r
}

func decodeCreateProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
  var req CreateUserRequest
  err = json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    return nil, err
  }
  return req, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
  var req = GetUserRequest{}
  vars := mux.Vars(r)
  req = GetUserRequest{
    Id: vars["id"],
  }
	return req, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
  var req = UpdateUserRequest{}
  if err := json.NewDecoder(r.Body).Decode(&req.user); err != nil {
      return nil, err
    }
  return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
  var req = DeleteUserRequest{}
  vars := mux.Vars(r)
  req = DeleteUserRequest{
    Id: vars["id"],
  }
  return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  return json.NewEncoder(w).Encode(response)
}
