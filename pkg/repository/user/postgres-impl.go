package user

import (
	"context"

	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/config/postgres"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/user"
)

type UserRepoImpl struct {
	pgCln postgres.PostgreClient
}

func NewUserRepo(pgcln postgres.PostgreClient) user.UserRepo {
	return &UserRepoImpl{pgCln: pgcln}
}

func (u *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (result user.User, err error) {
	db := u.pgCln.GetClient()
	db.Model(&user.User{}).First(&result, "email = ?", email)
	if err = db.Error; err != nil {
		return user.User{}, err
	}
	return result, err
}

func (u *UserRepoImpl) InsertUser(ctx context.Context, insertedUser *user.User) (err error) {
	db := u.pgCln.GetClient()
	result := db.Model(&user.User{}).Create(&insertedUser)
	if err := result.Error; err != nil {
		return err
	}
	return err
}
