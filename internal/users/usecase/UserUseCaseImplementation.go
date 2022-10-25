package usecase

import (
	"goWebServer/shared/domain/user"
	"time"
)

type UserUseCaseImplementation struct {
	UserPersistenceRepository user.UserPersistentRepository
	ContextTimeout            time.Duration
}

func NewUserUseCase(repo user.UserPersistentRepository, timeout time.Duration) user.UserUseCase {
	return &UserUseCaseImplementation{
		UserPersistenceRepository: repo,
		ContextTimeout:            timeout,
	}
}
