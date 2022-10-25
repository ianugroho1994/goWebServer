package user

import (
	"context"
)

// UserUseCase All representation of citizen use case
type UserUseCase interface {
	Fetch(ctx context.Context) (u []User, totalItem int64, err error)
	UpdateProfile(ctx context.Context, citizen *User) (u User, err error)
	GetByID(ctx context.Context, id string) (u User, err error)
}
