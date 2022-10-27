package admin

import (
	"goWebServer/shared/domain/user"
	"time"
)

type AdminDbModel struct {
	ID        int64      `json:"id" db:"id" gorm:"primaryKey"`
	UserId    int64      `json:"user_id" db:"user_id"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt *time.Time `json:"-" db:"updated_at"`
}

type Admin struct {
	ID        int64      `json:"id"`
	User      *user.User `json:"user" db:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}
