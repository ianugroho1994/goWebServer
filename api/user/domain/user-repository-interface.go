package domain

import "context"

type IUserRepository interface {
	GetList(ctx context.Context, cursor string, num int64) (result []User, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (user User, err error)
	GetByEmail(ctx context.Context, email string) (user User, err error)
	GetAll(ctx context.Context) (result []User, err error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
	ChangePassword(ctx context.Context, id int64, password string) error
}
