package controller

import (
	"net/http"
	"notification-api/service"

	"github.com/gin-gonic/gin"
)

type ResetPasswordController struct {
	Service service.ResetPasswordService
}

func NewResetPasswordController(service service.ResetPasswordService) ResetPasswordController {
	return ResetPasswordController{
		Service: service,
	}
}

func (rpc *ResetPasswordController) RequestReset(c *gin.Context) {
	var payload struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	rpc.Service.ResetPassword(payload.Email)

	c.JSON(http.StatusOK, gin.H{"message": "Token sent to email"})
}

func (rpc *ResetPasswordController) ValidateReset(c *gin.Context) {
	var payload struct {
		Email       string `json:"email"`
		Token       string `json:"token"`
		NewPassword string `json:"newPassword"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := rpc.Service.ValidUser(payload.Email, payload.Token, payload.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sucessfully updated password"})
}
