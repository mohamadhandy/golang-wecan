package user

import (
	"kitabisa/logger"

	"gorm.io/gorm"
)

type UserRepositoryDB interface {
	RegisterUser(User) (User, error)
	FindUserById(int) (User, error)
	FindUserByEmail(string) (User, error)
}

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) *userRepositoryDB {
	return &userRepositoryDB{db}
}

func (u *userRepositoryDB) RegisterUser(user User) (User, error) {
	var err error
	if err = u.db.Create(&user).Error; err != nil {
		logger.Error("Unexpected DB error!" + err.Error())
		return user, err
	}
	return user, nil
}

func (u *userRepositoryDB) FindUserByEmail(email string) (User, error) {
	var err error
	var user User
	if err = u.db.Where("email = ?", email).Find(&user).Error; err != nil {
		logger.Error("Unexpected DB error" + err.Error())
		return user, err
	}
	return user, nil
}

func (u *userRepositoryDB) FindUserById(id int) (User, error) {
	var err error
	var user User
	if err = u.db.Where("user_id ? = ", id).Find(&user).Error; err != nil {
		logger.Error("Unexpected DB error" + err.Error())
		return user, err
	}
	return user, nil
}
