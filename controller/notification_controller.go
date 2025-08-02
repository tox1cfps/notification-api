package controller

import (
	"net/http"
	"notification-api/model"
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
		var notification model.Notification
		err := c.ShouldBindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		insertedNotification, err := n.NotificationService.CreateNotification(notification)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, insertedNotification)
	}
}
