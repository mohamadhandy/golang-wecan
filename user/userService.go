package user

import (
	"errors"
	"kitabisa/logger"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(RegisterUserInput) (User, error)
	FindUserById(int) (User, error)
	Login(LoginInput) (User, error)
	SaveAvatar(int, string) (User, error)
}

type userService struct {
	userRepositoryDB UserRepositoryDB
}

func NewUserService(userRepo UserRepositoryDB) *userService {
	return &userService{userRepositoryDB: userRepo}
}

func (us *userService) RegisterUser(input RegisterUserInput) (User, error) {
	var user User
	user.Email = input.Email
	user.Occupation = input.Occupation
	user.Name = input.Name
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		logger.Error("Unexpected error " + err.Error())
	}
	user.PasswordHash = string(passwordHash)
	newUser, err := us.userRepositoryDB.RegisterUser(user)
	if err != nil {
		logger.Error("Unexpected Error: " + err.Error())
	}
	return newUser, nil
}

func (us *userService) FindUserById(id int) (User, error) {
	user, err := us.userRepositoryDB.FindUserById(id)
	if err != nil {
		logger.Error("Service error" + err.Error())
		return user, err
	}
	if user.ID == 0 {
		logger.Error("user not found")
		return user, errors.New("user not found")
	}
	return user, nil
}

func (us *userService) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := us.userRepositoryDB.FindUserByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (us *userService) SaveAvatar(userId int, fileLocation string) (User, error) {
	user, err := us.userRepositoryDB.FindUserById(userId)
	if err != nil {
		return user, err
	}
	user.AvatarFieldName = fileLocation

	updatedUser, err := us.userRepositoryDB.UpdateUser(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}
