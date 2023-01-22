package domain

import (
	"time"

	"smartlab/api/user/domain"
)

// Auth is used for authentication data inside persistent databases
type Authentication struct {
	ID           int64        `json:"id" db:"id"`
	UserID       int64        `json:"user_id" db:"user_id"`
	User         *domain.User `json:"user" db:"-"`
	AccessToken  string       `json:"access_token" db:"access_token"`
	TokenType    string       `json:"-" db:"token_type"`
	ExpiresIn    int64        `json:"expires_in" db:"expires_in"`
	RefreshToken string       `json:"refresh_token" db:"refresh_token"`
	IsBlacklist  bool         `json:"-" db:"is_blacklist"`
	XDeviceID    string       `json:"-" db:"x_device_id"`
	CreatedAt    time.Time    `json:"-" db:"created_at"`
	UpdatedAt    *time.Time   `json:"-" db:"updated_at"`
	// Jti          string       `json:"jti" db:"jti"`
}
