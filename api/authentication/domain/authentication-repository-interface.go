package domain

import "context"

type IAuthenticationRepository interface {
	Store(ctx context.Context, auth *Authentication) error
}
