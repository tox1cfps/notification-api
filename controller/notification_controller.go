package controller

import (
	"notification-api/service"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	NotificationService service.NotificationService
}

func NewNotificationController(service service.NotificationService) NotificationController {
	return NotificationController{
		NotificationService: service,
	}
}

func (n *NotificationController) Handle() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
