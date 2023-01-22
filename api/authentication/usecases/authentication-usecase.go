package usecases

import (
	"context"
	"hardtmann/api/authentication/domain"
	userDomain "hardtmann/smartlab/api/user/domain"
)

type AuthenticationUseCase struct {
}

func NewAuthenticationUseCase() domain.IAuthenticationUseCase {
	return &AuthenticationUseCase{}
}

func (a *AuthenticationUseCase) Login(ctx context.Context, user *userDomain.User, handsetID string) (authentication domain.Authentication, err error) {
	return domain.Authentication{}, nil
}
