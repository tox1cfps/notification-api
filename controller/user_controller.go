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

func (u *UserController) LoginUser(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
	}

	token, err := u.UserService.LoginUser(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
