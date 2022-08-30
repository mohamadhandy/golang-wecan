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
	token, err := u.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
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
	dtoUser := FormatUser("", newUser)
	response := helper.ResponseAPI(dtoUser, "success", http.StatusCreated, "Success Register User!")
	ctx.JSON(http.StatusCreated, response)
}

func (u *userHandler) UploadAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ResponseAPI(data, "error", http.StatusBadRequest, "Error upload file1")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := ctx.MustGet("currentUser").(User)
	userId := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ResponseAPI(data, "error", http.StatusBadRequest, "Error upload file2")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = u.userService.SaveAvatar(userId, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ResponseAPI(data, "error", http.StatusBadRequest, "Error upload file3")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.ResponseAPI(data, "success", http.StatusOK, "Success Upload Avatar")
	ctx.JSON(http.StatusOK, response)
}

func (u *userHandler) FindAllUser(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(User)
	if currentUser.ID != 0 && currentUser.Role == "Admin" {
		users, err := u.userService.GetAllByUser()
		if err != nil {
			response := helper.ResponseAPI(nil, "error", http.StatusBadRequest, "Error get All by user")
			ctx.JSON(http.StatusBadRequest, response)
			return
		} else {
			response := helper.ResponseAPI(users, "success", http.StatusOK, "Success get All User")
			ctx.JSON(http.StatusOK, response)
			return
		}
	} else {
		response := helper.ResponseAPI(nil, "error", http.StatusUnauthorized, "You dont have permissions")
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}
}
