package test

import (
	"testing"

	"usersvc/usersvc"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	email := "test@test.com"
	password := "12345"
	service, ctx := setup()
	createUser, err := service.CreateUser(ctx, email, password)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assert.Equal(t, "Success", createUser)
}

func TestGetUser(t *testing.T) {
	id := "5654abe9-85f9-43e6-8cb4-2159729fa109"
	service, ctx := setup()
	getUser, err := service.GetUser(ctx, id)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assert.Equal(t, "test@test.com", getUser)
}

func TestGetUserNotFound(t *testing.T) {
	service, ctx := setup()
	getUserNotFound, err := service.GetUser(ctx, "abc")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assert.Equal(t, getUserNotFound, "Not Found")
}

func TestGetUserNoID(t *testing.T) {
	service, ctx := setup()
	getUserNoID, err := service.GetUser(ctx, "")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assert.Equal(t, getUserNoID, "Repo need a correct ID")
}

func TestGetAllUsers(t *testing.T) {
	service, ctx := setup()
	getAllUsers, err := service.GetAllUsers(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assert.NotEmpty(t, getAllUsers)
}

func TestUpdateUser(t *testing.T) {
	user := usersvc.User{
		Id:       "5654abe9-85f9-43e6-8cb4-2159729fa109",
		Email:    "test2@test.com",
		Password: "12345",
	}
	service, ctx := setup()
	updateUser, err := service.UpdateUser(ctx, user)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	email := user.Email
	assert.Equal(t, "test2@test.com", email)
	assert.Equal(t, "successfully updated", updateUser)
}

func TestDeleteUser(t *testing.T) {
	id := "5654abe9-85f9-43e6-8cb4-2159729fa109"
	service, ctx := setup()
	deleteUser, err := service.DeleteUser(ctx, id)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assert.Equal(t, "Deleted", deleteUser)
}
