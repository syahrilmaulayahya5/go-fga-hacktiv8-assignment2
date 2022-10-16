package user

import (
	"context"
	"errors"

	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}

func (u *UserUsecaseImpl) GetUserByEmailSvc(ctx context.Context, email string) (result user.User, err error) {
	result, err = u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		err = errors.New("INTERNAL_SERVER_ERROR")
		return user.User{}, err
	}

	if result.ID <= 0 {
		err = errors.New("NOT_FOUND")
		return user.User{}, err
	}

	return result, err
}
