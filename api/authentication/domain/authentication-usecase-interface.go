package domain

import (
	"context"
	userDomain "hardtmann/api/user/domain"
)

type IAuthenticationUseCase interface {
	Login(ctx context.Context, user *userDomain.User, handsetID string) (authentication Authentication, err error)
}
