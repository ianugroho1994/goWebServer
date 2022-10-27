package mySql

import (
	"context"
	"fmt"
	"goWebServer/shared/domain/user"
	"goWebServer/shared/logger"

	"gorm.io/gorm"
)

type UserPersistenceRepositoryImplementation struct {
	Connection *gorm.DB
	tableName  string
}

func NewUserPersistenceRepository(Conn *gorm.DB) user.UserPersistentRepository {
	return &UserPersistenceRepositoryImplementation{
		Connection: Conn,
		tableName:  "users",
	}
}

func (impl UserPersistenceRepositoryImplementation) GetByID(ctx context.Context, id string) (userDomain user.User, err error) {
	res := impl.Connection.Table(impl.tableName).First(&userDomain, id)
	err = res.Error
	if err != nil {
		userDomain = user.User{}
		return
	}
	return
}

func (impl UserPersistenceRepositoryImplementation) Fetch(ctx context.Context) (userDomains []user.User, totalItem int64, err error) {
	res := impl.Connection.Table(impl.tableName).Find(&userDomains) //fetching queries into databases
	err = res.Error

	if err != nil {
		totalItem = 0
		userDomains = make([]user.User, 0)
		return
	}
	totalItem = int64(len(userDomains))
	return
}

func (impl UserPersistenceRepositoryImplementation) GetByEmail(ctx context.Context, email string) (userDomain user.User, err error) {
	res := impl.Connection.Table(impl.tableName).First(&userDomain, "email = ?", email) //fetching queries into databases
	err = res.Error

	if err != nil {
		userDomain = user.User{}
		return
	}
	return
}

func (impl UserPersistenceRepositoryImplementation) Update(ctx context.Context, user *user.User) (err error) {
	res := impl.Connection.Table(impl.tableName).Model(&user).Updates(map[string]interface{}{
		"name":          user.Name,
		"address":       user.Address,
		"starting_date": user.StartingDate,
		"email":         user.Email,
		"phone_number":  user.PhoneNumber,
	})
	return res.Error
}

func (impl UserPersistenceRepositoryImplementation) Store(ctx context.Context, user *user.User) (err error) {
	user.ID = "0"
	res := impl.Connection.Table(impl.tableName).Create(&user)
	err = res.Error
	affect := res.RowsAffected
	if res.Error != nil {
		return
	}

	if affect != 1 {
		err = fmt.Errorf("citizen created. total affected: %d", affect)
		logger.Log.Fatal(err)
		return
	}
	return
}

func (impl UserPersistenceRepositoryImplementation) Delete(ctx context.Context, id string) (err error) {
	res := impl.Connection.Table(impl.tableName).Where("id = ?", id).Delete(user.User{})
	err = res.Error
	affect := res.RowsAffected
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("weird  behavior. total affected: %d", affect)
		logger.Log.Fatal(err)
		return
	}
	return
}
