package user

import (
	"context"
	_ "time"
)

// UserPersistentRepository UserRepository represent the user's repository contract
type UserPersistentRepository interface {
	GetByID(ctx context.Context, id string) (User, error)
	Fetch(ctx context.Context) (res []User, totalItem int64, err error)
	GetByEmail(ctx context.Context, email string) (u User, err error)
	Update(ctx context.Context, user *User) (err error)
	Store(ctx context.Context, user *User) (err error)
	Delete(ctx context.Context, id string) (err error)
}
