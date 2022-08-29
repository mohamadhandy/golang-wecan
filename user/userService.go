package user

import (
	"kitabisa/logger"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(RegisterUserInput) (User, error)
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
