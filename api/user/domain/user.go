package domain

type User struct {
	ID                   int64  `json:"id" db:"id"`
	Username             string `json:"username" db:"username"`
	Email                string `json:"email" db:"email"`
	Password             string `json:"password,omitempty" db:"password"`
	ResetPasswordToken   string `json:"-" db:"reset_password_token"`
	ResetPasswordExpired int64  `json:"-" db:"reset_password_expired"`
}
