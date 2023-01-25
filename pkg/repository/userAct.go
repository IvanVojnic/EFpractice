package repository

import (
	"EFpractic2/models"
	"github.com/jmoiron/sqlx"
)

type UserActPostgres struct {
	db *sqlx.DB
}

func NewUserActPostgres(db *sqlx.DB) *UserActPostgres {
	return &UserActPostgres{db: db}
}

func (r *UserActPostgres) CreateUser(user models.User) error {
	//return s.repo.CreateUser(user)
	return nil
}

func (r *UserActPostgres) UpdateUser(user models.User) error {
	//return s.repo.UpdateUser(user)
	return nil
}

func (r *UserActPostgres) GetUser(userId int) (models.User, error) {
	//return s.repo.GetUser(userId)
	user := models.User{UserId: 1, UserAge: 20, UserIsRegular: true, UserName: "Jhon"}
	return user, nil
}

func (r *UserActPostgres) DeleteUser(userId int) error {
	//return s.repo.DeleteUser(userId)
	return nil
}

func (r *UserActPostgres) GetAllUsers() ([]models.User, error) {
	//return s.repo.GetAllUsers()
	var user1 = models.User{UserId: 1, UserAge: 20, UserIsRegular: true, UserName: "Jon"}
	var user2 = models.User{UserId: 2, UserAge: 21, UserIsRegular: true, UserName: "Jack"}
	users := make([]models.User, 0)
	users = append(users, user1)
	users = append(users, user2)
	return users, nil
}
