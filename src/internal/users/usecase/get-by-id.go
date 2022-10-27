package usecase

import (
	"context"
	"goWebServer/shared/domain/user"
	"goWebServer/shared/logger"
)

func (impl UserUseCaseImplementation) GetByID(ctx context.Context, id string) (u user.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, impl.ContextTimeout)
	defer cancel()
	logger.Log.Info("")
	u, err = impl.UserPersistenceRepository.GetByID(ctx, id)
	return
}
