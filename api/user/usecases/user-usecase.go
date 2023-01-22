package usecases

import (
	"context"
	"fmt"
	"hardtmann/smartlab/api/user/domain"
	"log"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo           domain.IUserRepository
	ContextTimeout time.Duration
}

func NewUserUseCase(userRepo domain.IUserRepository, timeout time.Duration) domain.IUserUseCase {
	return &UserUseCase{
		repo:           userRepo,
		ContextTimeout: timeout,
	}
}

func (u *UserUseCase) ChangePassword(ctx context.Context, id int64, newPassword string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()
	saltedPassword := fmt.Sprintf("%s.%s", newPassword, viper.GetString("password_salt"))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}
	err = u.repo.ChangePassword(ctx, id, string(hashedPassword))

	return err
}
