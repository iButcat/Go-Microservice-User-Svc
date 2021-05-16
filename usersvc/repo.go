package usersvc

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"gorm.io/gorm"
)

type repo struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepo(db *gorm.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) (string, error) {
	if user.Email == "" || user.Password == "" {
		return "Empty", nil
	}

	if err := repo.db.AutoMigrate(&user); err != nil {
		return "Can't migrate", err
	}

	if err := repo.db.Create(&user).Error; err != nil {
		return "", err
	}
	return user.Id + "successfully created", nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	var user = new(User)
	if len(id) > 0 {
		if err := repo.db.Table("users").Where(
			"id = ?", id).Scan(&user).Error; err != nil {
			errors.Is(err, gorm.ErrRecordNotFound)
			return "Not Found", err
		}
	} else {
		return "Repo need a correct ID", nil
	}
	email = user.Email
	return email, nil
}

func (repo *repo) GetAllUsers(ctx context.Context) ([]User, error) {
	var allUsers []User
	if err := repo.db.Raw(
		"SELECT * FROM users").Scan(&allUsers).Error; err != nil {
		return nil, err
	}
	return allUsers, nil
}

func (repo *repo) UpdateUser(ctx context.Context, user User) (string, error) {
	if err := repo.db.Exec(
		"UPDATE users SET email=?, password=? WHERE id = ?",
		user.Email,
		user.Password,
		user.Id).Error; err != nil {
		return "", err
	}
	return user.Id + "successfully updated", nil
}

func (repo *repo) DeleteUser(ctx context.Context, id string) (string, error) {
	if err := repo.db.Exec("DELETE FROM users WHERE id= ?", id).Error; err != nil {
		return "", err
	}
	return "Succesfully deleted", nil
}
