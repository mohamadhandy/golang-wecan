package user

import (
	"kitabisa/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService userService
}

func NewUserHandler(userService userService) *userHandler {
	return &userHandler{userService: userService}
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
