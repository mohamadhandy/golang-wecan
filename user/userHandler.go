package user

import (
	"fmt"
	"kitabisa/auth"
	"kitabisa/helper"
	"kitabisa/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService userService
	authService auth.Service
}

func NewUserHandler(userService userService, authService auth.Service) *userHandler {
	return &userHandler{userService: userService, authService: authService}
}

func (u *userHandler) Login(ctx *gin.Context) {
	var input LoginInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ResponseAPI(nil, "error", http.StatusUnprocessableEntity, "Login account failed")
		logger.Error("error" + err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedinUser, err := u.userService.Login(input)
	if err != nil {
		response := helper.ResponseAPI(nil, "error", http.StatusUnprocessableEntity, "Login account failed")
		logger.Error("error" + err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
	}
	fmt.Println("Loggedin user id", loggedinUser.ID)
	fmt.Println("Loggedin user id", loggedinUser)
	token, err := u.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		fmt.Println("TEST ERROR" + err.Error())
		response := helper.ResponseAPI(nil, "error", http.StatusBadRequest, "Login failed")
		logger.Error("error" + err.Error())
		ctx.JSON(http.StatusBadRequest, response)
	}
	dtoLoggedinUser := FormatUser(token, loggedinUser)
	response := helper.ResponseAPI(dtoLoggedinUser, "success", http.StatusOK, "Success Login")
	ctx.JSON(http.StatusOK, response)
}

func (u *userHandler) RegisterUser(ctx *gin.Context) {
	var input RegisterUserInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ResponseAPI(nil, "error", http.StatusUnprocessableEntity, "Register account failed")
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := u.userService.RegisterUser(input)
	if err != nil {
		response := helper.ResponseAPI(nil, "error", http.StatusBadRequest, "Register account failed")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	// generate token, and forced token to db
	// token, err := u.authService.GenerateToken(newUser.ID)
	// logger.Info(token)
	// if err != nil {
	// 	response := helper.ResponseAPI(nil, "error", http.StatusBadRequest, "Register account failed")
	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }
	dtoUser := FormatUser("", newUser)
	response := helper.ResponseAPI(dtoUser, "success", http.StatusCreated, "Success Register User!")
	ctx.JSON(http.StatusCreated, response)
}
