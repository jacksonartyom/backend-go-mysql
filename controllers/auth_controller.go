package controllers

import (
	"backend-go-mysql/dto/request"
	"backend-go-mysql/dto/response"
	"backend-go-mysql/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
}

func (a *AuthController) Login(c *gin.Context) {
	var input request.UserDto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := a.AuthService.Login(input.Email, input.Password)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	res := response.SuccessResponse[response.UserResponse]{
		Message: "success",
		Result:  user,
	}

	c.JSON(200, res)
}
