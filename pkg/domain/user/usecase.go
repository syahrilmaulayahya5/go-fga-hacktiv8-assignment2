package user

import (
	"context"
)

type UserUsecase interface {
	GetUserByEmailSvc(ctx context.Context, email string) (result User, err error)
}
