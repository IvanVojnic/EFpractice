package repository

import (
	"EFpractic2/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserActPostgres struct {
	db *pgxpool.Pool
}

func NewUserActPostgres(db *pgxpool.Pool) *UserActPostgres {
	return &UserActPostgres{db: db}
}

func (r *UserActPostgres) CreateUser(ctx context.Context, user models.User) error {
	err := r.db.QueryRow(ctx, "insert into user (name, age, isRegular, password) select $1, $2, $3, $4, $5",
		user.UserName, user.UserAge, user.UserIsRegular, user.Password)
	if err != nil {
		return fmt.Errorf("Error while user creating: %v", err)
	}
	return nil
}

func (r *UserActPostgres) UpdateUser(ctx context.Context, user models.User) error {
	res, err := r.db.Exec(ctx, "UPDATE user SET name = $1, age = $2, isRegular =$3 WHERE id = $4", user.UserName, user.UserAge, user.UserIsRegular, user.UserId)
	if err != nil {
		return err
	}
	fmt.Printf("[update a row] updated num rows: %d", res.RowsAffected())
	return nil
}

func (r *UserActPostgres) GetUser(ctx context.Context, userId int) (models.User, error) {
	//return s.repo.GetUser(userId)
	user := models.User{}
	err := r.db.QueryRow(ctx, "select * from user where id=$1 and not deleted", userId).Scan(
		&user.UserId, &user.UserName, &user.UserAge, &user.Password, &user.UserIsRegular)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserActPostgres) DeleteUser(ctx context.Context, userId int) error {
	commandTag, err := r.db.Exec(context.Background(), "delete from user where id=$1", userId)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return errors.New("No row found to delete")
	}
	return nil
}

func (r *UserActPostgres) GetAllUsers(ctx context.Context) ([]models.User, error) {
	//return s.repo.GetAllUsers()
	var user1 = models.User{UserId: 1, UserAge: 20, UserIsRegular: true, UserName: "Jon"}
	var user2 = models.User{UserId: 2, UserAge: 21, UserIsRegular: true, UserName: "Jack"}
	users := make([]models.User, 0)
	users = append(users, user1)
	users = append(users, user2)
	return users, nil
}
