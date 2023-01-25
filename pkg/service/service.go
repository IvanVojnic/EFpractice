package service

import (
	"EFpractic2/models"
	"EFpractic2/pkg/repository"
)

type UserAct interface {
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	GetUser(int) (models.User, error)
	DeleteUser(int) error
	GetAllUsers() ([]models.User, error)
}

type Service struct {
	UserAct
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserAct: NewUserActSrv(repos.UserAct),
	}
}
