package service

import (
	"EFpractic2/models"
	"EFpractic2/pkg/repository"
	"context"
)

type UserAct interface {
	CreateUser(context.Context, models.User) error
	UpdateUser(context.Context, models.User) error
	GetUser(context.Context, int) (models.User, error)
	DeleteUser(context.Context, int) error
	GetAllUsers(context.Context) ([]models.User, error)
}

type Service struct {
	UserAct
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserAct: NewUserActSrv(repos.UserAct),
	}
}
