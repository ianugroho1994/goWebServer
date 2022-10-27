package usecase

import (
	"context"
	"goWebServer/shared/domain/user"
)

func (impl UserUseCaseImplementation) UpdateProfile(ctx context.Context, citizen *user.User) (u user.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, impl.ContextTimeout)
	defer cancel()

	err = impl.UserPersistenceRepository.Update(ctx, citizen)
	return
}
