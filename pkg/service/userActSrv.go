package service

import (
	"EFpractic2/models"
	"EFpractic2/pkg/repository"
)

type UserActSrv struct {
	repo repository.UserAct
}

func NewUserActSrv(repo repository.UserAct) *UserActSrv {
	return &UserActSrv{repo: repo}
}

func (s *UserActSrv) CreateUser(user models.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserActSrv) UpdateUser(user models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserActSrv) GetUser(userId int) (models.User, error) {
	return s.repo.GetUser(userId)
}

func (s *UserActSrv) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}

func (s *UserActSrv) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}
