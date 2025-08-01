package main

import (
	"notification-api/client"
	"notification-api/config"
	"notification-api/controller"
	"notification-api/db"
	"notification-api/repository"
	"notification-api/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	config.Init()
	setting := config.Settings
	dbconnection, err := db.ConnectDB(
		setting.Database.Host,
		setting.Database.Port,
		setting.Database.User,
		setting.Database.Password,
		setting.Database.DbName,
		setting.Database.SslMode,
	)
	if err != nil {
		panic(err)
	}

	NotificationRepository := repository.NewNotificationRepository(dbconnection)

	Gmailsmtp := client.NewGmailsmtpClient()

	NotificationService := service.NewNotificationService(NotificationRepository, Gmailsmtp)

	NotificationController := controller.NewNotificationController(NotificationService)

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Pong",
		})
	})
	r.POST("/sendEmail", NotificationController.Handle())
	r.Run(":8080")
}
