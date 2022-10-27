package admin

import (
	"context"
)

// AdminUseCase AdminUseCase represent the admin's use case contract
type AdminUseCase interface {
	GetByID(ctx context.Context, ID int64) (adminDomain Admin, err error)
	GetProfile(ctx context.Context, email string) (adminDomain Admin, err error)
	Update(ctx context.Context, admin *Admin) (err error)
	UpdateProfile(ctx context.Context, admin *Admin) (err error)
	Fetch(ctx context.Context, offset int64, num int64, search string) (res []Admin, total_item int64, err error)
	Store(ctx context.Context, admin *Admin) (err error)
	Delete(ctx context.Context, id int64) (statusCode int64, err error)
	RegisterAdmin(ctx context.Context, admin *Admin) (err error, statusCode int64)
}
