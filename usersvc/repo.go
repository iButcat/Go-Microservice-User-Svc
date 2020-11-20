package usersvc

import (
	"context"

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
	_ = repo.db.AutoMigrate(&user)

	if err := repo.db.Create(&user).Error; err != nil {
		return "", err
	}
	return "successfully created", nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	if err := repo.db.Raw("SELECT email FROM users WHERE id=id", id).Scan(&email).Error; err != nil {
		return "Not Found", err
	}
	return email, nil
}

func (repo *repo) UpdateUser(ctx context.Context, user User) (string, error) {
	if err := repo.db.Exec("UPDATE users SET email=?, password=? WHERE id = ?", user.Email, user.Password, user.Id).Error; err != nil {
		return "", err
	}
	return "successfully updated", nil
}

func (repo *repo) DeleteUser(ctx context.Context, id string) (string, error) {
	if err := repo.db.Exec("DELETE FROM users WHERE id= ?", id).Error; err != nil {
		return "", err
	}
	return "Succesfully deleted", nil
}
