package controller

import (
	"net/http"
	"notification-api/model"
	"notification-api/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) UserController {
	return UserController{
		UserService: service,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	insertedUser, err := u.UserService.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusAccepted, insertedUser)
}
