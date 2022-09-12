package app

import (
	"fmt"
	"kitabisa/auth"
	"kitabisa/campaign"
	"kitabisa/helper"
	"kitabisa/logger"
	"kitabisa/transaction"
	"kitabisa/user"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error when loading env")
	} else {
		logger.Info("Run load env file smoothly")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	serverPort := os.Getenv("SERVER_PORT")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	// initialize repositoryDB
	userRepo := user.NewUserRepositoryDB(db)
	campaignRepo := campaign.NewCampaignRepositoryDB(db)
	transactionRepo := transaction.NewTransactionRepositoryDB(db)

	// initialize service
	userService := user.NewUserService(userRepo)
	authService := auth.NewService()
	campaignService := campaign.NewCampaignService(campaignRepo)
	transactionService := transaction.NewTransactionService(*transactionRepo)

	// initialize handler
	userHandler := user.NewUserHandler(*userService, authService)
	campaignHandler := campaign.NewCampaignHandler(campaignService, authService)
	transactionHandler := transaction.NewTransactionHandler(transactionService)

	if err != nil {
		logger.Fatal("Error connection")
	}

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.GET("/users", authMiddleware(authService, userService), userHandler.FindAllUser)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	// CAMPAIGNS
	api.GET("/campaigns", authMiddleware(authService, userService), campaignHandler.FindAllCampaigns)
	api.GET("/campaigns/:campaignid", authMiddleware(authService, userService), campaignHandler.FindCampaignById)
	api.PUT("/campaigns/:campaignid", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)

	api.GET("/campaigns/:campaignid/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	routerRun := fmt.Sprintf(":%s", serverPort)
	router.Run(routerRun)
}

func authMiddleware(authService auth.Service, userService user.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ResponseAPI(nil, "error", http.StatusUnauthorized, "Unauthorized user!")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// Bearer Token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		ok, userId, err := authService.ValidateToken(tokenString)
		if err != nil && !ok && userId == -1 {
			response := helper.ResponseAPI(nil, "error", http.StatusUnauthorized, "Unauthorized user!")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		} else {
			user, err := userService.FindUserById(userId)
			if err != nil {
				response := helper.ResponseAPI(nil, "error", http.StatusUnauthorized, "Unauthorized user")
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			ctx.Set("currentUser", user)
		}
	}
}
