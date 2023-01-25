package repository

import (
	"EFpractic2/models"
	"github.com/jmoiron/sqlx"
)

type UserAct interface {
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	GetUser(int) (models.User, error)
	DeleteUser(int) error
	GetAllUsers() ([]models.User, error)
}

type Repository struct {
	UserAct
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserAct: NewUserActPostgres(db),
	}
}
