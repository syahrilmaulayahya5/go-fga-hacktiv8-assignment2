package user

import "context"

type UserRepo interface {
	GetUserByEmail(ctx context.Context, email string) (result User, err error)
	InsertUser(ctx context.Context, user *User) (err error)
}
