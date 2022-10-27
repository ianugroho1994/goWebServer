package admin

import (
	"context"
)

// AdminPersistentRepository AdminRepository represent the admin's repository contract
type AdminPersistentRepository interface {
	Fetch(ctx context.Context, offset int64, num int64, search string) (res []Admin, total_item int64, err error)
	GetByID(ctx context.Context, id int64) (adminDomain Admin, err error)
	GetByUserID(ctx context.Context, id int64) (adminDomain Admin, err error)
	GetByEmail(ctx context.Context, email string) (adminDomain Admin, err error)
	Update(ctx context.Context, ar *Admin) (err error)
	Store(ctx context.Context, a *Admin) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
