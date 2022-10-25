package usecase

import (
	"context"
	"goWebServer/shared/domain/user"
	"goWebServer/shared/logger"
)

func (impl UserUseCaseImplementation) Fetch(ctx context.Context) (u []user.User, totalItem int64, err error) {
	ctx, cancel := context.WithTimeout(ctx, impl.ContextTimeout)
	defer cancel()
	logger.Log.Info("")
	u, totalItem, err = impl.UserPersistenceRepository.Fetch(ctx)
	if err != nil {
		totalItem = 0
		logger.Log.Fatal(err.Error())
		return
	}
	return
}
