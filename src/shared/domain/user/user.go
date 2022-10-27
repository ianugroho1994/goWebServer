package user

import (
	"fmt"
	"time"

	"goWebServer/shared/logger"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

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

func (user User) ChangePassword(newPassword string) (newUser User, err error) {
	saltedPassword := fmt.Sprintf("%s.%s", newPassword, viper.GetString("password_salt"))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("%s", "hashing password error")
		logger.Log.Error(err)
		return
	}
	newUser.Password = string(hashedPassword)
	return
}
