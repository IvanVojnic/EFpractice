package service

import (
	"EFpractic2/models"
	"EFpractic2/pkg/repository"
	"context"
)

type UserActSrv struct {
	repo repository.UserAct
}

func NewUserActSrv(repo repository.UserAct) *UserActSrv {
	return &UserActSrv{repo: repo}
}

func (s *UserActSrv) CreateUser(ctx context.Context, user models.User) error {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserActSrv) UpdateUser(ctx context.Context, user models.User) error {
	return s.repo.UpdateUser(ctx, user)
}

func (s *UserActSrv) GetUser(ctx context.Context, userId int) (models.User, error) {
	return s.repo.GetUser(ctx, userId)
}

func (s *UserActSrv) DeleteUser(ctx context.Context, userId int) error {
	return s.repo.DeleteUser(ctx, userId)
}

func (s *UserActSrv) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllUsers(ctx)
}
