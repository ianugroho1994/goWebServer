package repository

import (
	"context"
	"fmt"
	"hardtmann/smartlab/api/user/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	database  *gorm.DB
	tableName string
}

func NewUserRepository(db *gorm.DB) domain.IUserRepository {
	return &UserRepository{
		database:  db,
		tableName: "users",
	}
}

func (u *UserRepository) GetList(ctx context.Context, cursor string, num int64) (result []domain.User, nextCursor string, err error) {
	return nil, "", nil
}

func (u *UserRepository) GetByID(ctx context.Context, id int64) (user domain.User, err error) {
	res := u.database.Table(u.tableName).First(&user, id)

	err = res.Error
	if err != nil {
		user = domain.User{}
		return
	}
	return
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (user domain.User, err error) {
	res := u.database.Table(u.tableName).First(&user, "email = ?", email) //fetching queries into databases

	err = res.Error
	if err != nil {
		user = domain.User{}
		return
	}
	return
}

func (u *UserRepository) GetAll(ctx context.Context) (result []domain.User, err error) {
	res := u.database.Table(u.tableName).Find(&result) //fetching queries into databases
	err = res.Error

	if err != nil {
		result = make([]domain.User, 0)
		return
	}
	return
}

func (u *UserRepository) Create(ctx context.Context, user *domain.User) error {
	user.ID = int64(0)

	res := u.database.Table(u.tableName).Create(&user)
	err := res.Error

	affect := res.RowsAffected
	if res.Error != nil {
		return err
	}

	if affect != 1 {
		return err
	}
	return nil
}

func (u *UserRepository) Update(ctx context.Context, user *domain.User) error {
	res := u.database.Table(u.tableName).Model(&user).Updates(map[string]interface{}{
		"email": user.Email,
	})
	return res.Error
}

func (u *UserRepository) Delete(ctx context.Context, id int64) error {
	res := u.database.Table(u.tableName).Where("id = ?", id).Delete(domain.User{})

	err := res.Error
	affect := res.RowsAffected
	if err != nil {
		return err
	}

	if affect != 1 {
		return err
	}
	return nil
}

func (u *UserRepository) ChangePassword(ctx context.Context, id int64, password string) error {
	res := u.database.Table(u.tableName).Where("id = ?", id).Updates(map[string]interface{}{
		"password": password,
	})

	err := res.Error
	if err != nil {
		return err
	}

	affected := res.RowsAffected
	if affected > 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affected)
		return err
	}

	return nil
}
