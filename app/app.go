package app

import (
	"fmt"
	"kitabisa/logger"
	"os"

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

	fmt.Println("db", db)

	if err != nil {
		logger.Fatal("Error connection")
	}

	router := gin.Default()

	api := router.Group("/api/v1")
	fmt.Println("api", api)
	routerRun := fmt.Sprintf(":%s", serverPort)
	router.Run(routerRun)
}
