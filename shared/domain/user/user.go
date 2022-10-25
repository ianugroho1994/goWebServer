package user

import "time"

// User ...
type User struct {
	ID           string     `json:"id" db:"id" gorm:"primaryKey"`
	Name         string     `json:"name" db:"name"`
	Address      string     `json:"address" db:"address"`
	StartingDate time.Time  `json:"starting_date" db:"starting_date"`
	Email        string     `json:"email" db:"email"`
	PhoneNumber  string     `json:"phone_number" db:"phone_number"`
	Password     string     `json:"password" db:"password"`
	CreatedAt    time.Time  `json:"-" db:"created_at"`
	UpdatedAt    *time.Time `json:"-" db:"updated_at"`
}
