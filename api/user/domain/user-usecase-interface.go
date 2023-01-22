package domain

import "context"

type IUserUseCase interface {
	ChangePassword(ctx context.Context, id int64, newPassword string) error
}
